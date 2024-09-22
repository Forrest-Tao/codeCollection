package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"sync"
	"time"
)

var (
	etcdClient *clientv3.Client
	err        error
	wg         sync.WaitGroup
	session    *concurrency.Session
	election   *concurrency.Election
)

func main() {
	connectToETCD()
	defer etcdClient.Close()

	//wg.Add(1)
	//go func() {
	//	watchKey("key-1")
	//}()
	//putKeyValue("key-1", "value-1")
	//time.Sleep(time.Second)
	//
	//putKeyValue("key-1", "value-n")
	//time.Sleep(time.Second)
	//
	//lockKey("key-1")
	//time.Sleep(time.Second)
	//
	//deleteKey("key-1")
	//time.Sleep(time.Second)
	//
	//wg.Done()
	//wg.Wait()

	elect()
}

func connectToETCD() {
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 节点地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("无法连接到 etcd: %v", err)
	}

	fmt.Println("成功连接到 etcd！")
}

func putKeyValue(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = etcdClient.Put(ctx, key, value)
	if err != nil {
		log.Fatalf("无法创建键值对: %v", err)
	}
	fmt.Println("键值对创建成功！", key, value)
}

func deleteKey(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = etcdClient.Delete(ctx, key)
	if err != nil {
		log.Fatalf("无法删除键值对: %v", err)
	}
	fmt.Println("键值对删除成功！", key)
}

func getKeyValue(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := etcdClient.Get(ctx, key)
	if err != nil {
		return "", err
	}
	//返回的是一个数组 Kvs []*mvccpb.KeyValue,这里取第一个valye即可
	return string(resp.Kvs[0].Value), nil
}

func watchKey(key string) {
	watchChan := etcdClient.Watch(context.Background(), key)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Printf("事件类型: %s, 键: %s, 值: %s\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}

func lockKey(key string) {
	session, err = concurrency.NewSession(etcdClient)
	if err != nil {
		log.Fatalf("创建 session 失败: %v", err)
	}
	defer session.Close()

	// 创建一个分布式锁
	mutex := concurrency.NewMutex(session, "/my-lock/")

	// 获取锁
	if err = mutex.Lock(context.Background()); err != nil {
		log.Fatalf("获取锁失败: %v", err)
	}
	fmt.Println("获取锁成功")

	// 模拟临界区代码
	time.Sleep(2 * time.Second)

	// 释放锁
	if err = mutex.Unlock(context.Background()); err != nil {
		log.Fatalf("释放锁失败: %v", err)
	}
	fmt.Println("释放锁成功")
}

func elect() {
	session, err = concurrency.NewSession(etcdClient)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	election = concurrency.NewElection(session, "/my-election/")

	// 参加选举
	if election.Campaign(context.TODO(), "my-leader") != nil {
		log.Fatalf("竞选失败: %v", err)
	}

	// 当前实例成为领导者
	fmt.Println("获得领导权")

	// 等待一段时间
	time.Sleep(10 * time.Second)

	// 释放领导权
	if election.Resign(context.Background()) != nil {
		log.Fatalf("释放领导权失败: %v", err)
	}
}
