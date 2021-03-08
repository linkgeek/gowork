package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err error
	kv clientv3.KV
	putResp *clientv3.PutResponse
	getResp *clientv3.GetResponse
	delResp *clientv3.DeleteResponse
	kvpair *mvccpb.KeyValue
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

	kv = clientv3.NewKV(client)
}

// put
func etcdPut()  {
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job2", "hello job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("PrevVal: ", string(putResp.PrevKv.Value))
		}
	}
}

// get
func etcdGet()  {
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job2"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetResponse: ", getResp.Kvs)
	}
}

// del
func etcdDel()  {
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/",clientv3.WithPrefix(), clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(delResp.Deleted)

	if (len(delResp.PrevKvs)) != 0 {
		for _, kvpair = range delResp.PrevKvs {
			fmt.Println("删除了：", string(kvpair.Key), string(kvpair.Value))
		}
	}

}

func main()  {
	//etcdPut()
	//etcdGet()
	etcdDel()
	//etcdGet()
}
