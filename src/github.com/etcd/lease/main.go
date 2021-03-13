package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err error
	lease clientv3.Lease
	leaseGrantResp *clientv3.LeaseGrantResponse
	leaseId clientv3.LeaseID
	kv clientv3.KV
	putResp *clientv3.PutResponse
	getResp *clientv3.GetResponse
	keepResp *clientv3.LeaseKeepAliveResponse
	keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
)

func init()  {
	config = clientv3.Config{
		Endpoints:   []string{"192.168.7.26:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 获取kv api子集
	kv = clientv3.NewKV(client)

	// 申请租约
	lease = clientv3.NewLease(client)

	// 申请一个10秒租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约ID
	leaseId = leaseGrantResp.ID

	// 自动续约
	//ctx,_:= context.WithTimeout(context.TODO(), 5 * time.Second)

	// 5秒后取消自动续约
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp = <- keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已失效")
					goto END
				} else {
					fmt.Println("收到自动续约应答：", keepResp.ID)
				}
			}
		}
		END:
	}()

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job3", "job3", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功：", putResp.Header.Revision)

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job3"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了...")
			break
		}
		fmt.Println("没过期: ", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}

func main()  {

}
