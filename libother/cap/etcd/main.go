package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

var (
	cli *clientv3.Client
	err error
)

func init() {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:32770", "localhost:32771", "localhost:32772"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("connect to etcd success")
}

func main() {
	defer cli.Close()
	//PutAndGet()
	//Watch()
	//Lease()
	DistributedLock()
}

func PutAndGet() {
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "hello", "world")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "hello")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s6:%s6\n", ev.Key, ev.Value)
	}
}

func Watch() {
	// watch key:lmh change
	rch := cli.Watch(context.Background(), "hello") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s6 Key:%s6 Value:%s6\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func Lease() {
	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 2)
	if err != nil {
		log.Fatal(err)
	}

	// 10秒钟之后, /lmh/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/ttt/", "许磊", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("授权租约成功")
}

func KeepALive() {
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

func DistributedLock() {
	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}

	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}
