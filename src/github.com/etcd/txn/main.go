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
	keepResp *clientv3.LeaseKeepAliveResponse
	keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
	ctx context.Context
	cancelFunc context.CancelFunc
	txn clientv3.Txn
	txnResp *clientv3.TxnResponse
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

	// 3个步骤：
	// lease实现锁自动过期:
	// op操作
	// txn事务: if else then

	// 1, 上锁 (创建租约, 自动续租, 拿着租约去抢占一个key)
	lease = clientv3.NewLease(client)

	// 申请一个5秒租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约ID
	leaseId = leaseGrantResp.ID

	// 准备一个用于取消自动续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())

	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	// 5秒后取消自动续约
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp = <- keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已失效了")
					goto END
				} else {
					fmt.Println("收到自动续约应答：", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 获取kv api子集
	kv = clientv3.NewKV(client)

	// 创建事务
	txn = kv.Txn(context.TODO())

	// 定义事务
	// if 不存在key， then 设置它, else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "job99", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9"))

	// 提交事务
	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	// 判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用：", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 2, 业务处理
	fmt.Println("业务处理...")
	time.Sleep(5 * time.Second)

	// 3, 释放锁(取消自动续租, 释放租约)
	// defer 会把租约释放掉, 关联的KV就被删除了
}

func main()  {

}
