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
	kv clientv3.KV
	putOp clientv3.Op
	opResp clientv3.OpResponse
	getOp clientv3.Op
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

	// put
	putOp = clientv3.OpPut("/cron/jobs/job8", "88888")
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision:", opResp.Put().Header.Revision)

	// get
	getOp = clientv3.OpGet("/cron/jobs/job8")
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("获取Revision:", opResp.Get().Header.Revision)
	fmt.Println("获取Value:", string(opResp.Get().Kvs[0].Value))
}

func main()  {

}
