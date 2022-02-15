package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// #################################
// Redis.
// #################################

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func RedisHelloWorld() {
	// step: 0 - connect.
	redisClient := NewRedisClient()

	// "ping" Redis
	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	// step: 1 - write.
	err = redisClient.Set("name", "redis is awesome", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// step: 2 - read.
	val, err := redisClient.Get("name").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}

func RedisTimeExecSet(cnt int) time.Duration {
	start := time.Now()

	cl := NewRedisClient()

	for i := 0; i < cnt; i++ {
		if err := cl.Set("name", "redis is awesome", 0).Err(); err != nil {
			log.Fatal(err)
		}
	}

	return time.Since(start)
}

func RedisTimeExecGet(cnt int) time.Duration {
	start := time.Now()

	cl := NewRedisClient()

	for i := 0; i < cnt; i++ {
		if err := cl.Get("name").Err(); err != nil {
			log.Fatal(err)
		}
	}

	return time.Since(start)
}

// #################################
// etcd.
// #################################

func EtcdHelloWorld() {
	ctx := context.Background()

	// step: 0 - connect
	cl, err := NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if er := cl.Close(); er != nil {
			log.Fatal(er)
		}
	}()

	// step: 1 - write.
	_, err = cl.Put(ctx, "name2", "etcd is great")
	if err != nil {
		log.Fatal(err)
	}

	// step: 2 - read.
	res, er := cl.Get(ctx, "name2")
	if er != nil {
		log.Fatal(er)
	}
	fmt.Printf("res.Kvs[0].Version: %+v\n", res.Kvs[0].Version)
	fmt.Printf("res.Kvs[0].Value: %+s\n", res.Kvs[0].Value)
}

func NewEtcdClient() (*clientv3.Client, error) {
	cl, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("new etcd client err: %w", err)
	}

	return cl, nil
}

func EtcdTimeExecSet(cnt int) time.Duration {
	ctx := context.Background()
	start := time.Now()

	cl, err := NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if er := cl.Close(); er != nil {
			log.Fatal(er)
		}
	}()

	for i := 0; i < cnt; i++ {
		_, err = cl.Put(ctx, "name", "etcd_value_"+fmt.Sprintf("%d", i))
		if err != nil {
			log.Fatal(err)
		}
	}

	return time.Since(start)
}

func EtcdTimeExecGet(cnt int) time.Duration {
	ctx := context.Background()
	start := time.Now()

	cl, err := NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if er := cl.Close(); er != nil {
			log.Fatal(er)
		}
	}()

	for i := 0; i < cnt; i++ {
		_, err = cl.Get(ctx, "name")
		if err != nil {
			log.Fatal(err)
		}
	}
	return time.Since(start)
}

func main() {
	fmt.Println("Redis 'hello world'")
	RedisHelloWorld()

	fmt.Println("\netcd 'hello world'")
	EtcdHelloWorld()

	const count = 10_000

	fmt.Println("\nTime execution")
	fmt.Printf("redis set: %+v\n", RedisTimeExecSet(count)) // 7.825733292s
	fmt.Printf("redis get: %+v\n", RedisTimeExecGet(count)) // 7.625631125s

	fmt.Printf("etcd set: %+v\n", EtcdTimeExecSet(count)) // 4m21.974348792s
	fmt.Printf("etcd get: %+v\n", EtcdTimeExecGet(count)) // 1.302687666s

}
