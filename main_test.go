package main

import (
	"context"
	"fmt"
	"log"
	"testing"
)

// go test -bench=. -benchmem
// go test -bench=Redis -benchmem
// go test -bench=Etcd -benchmem

func BenchmarkRedisSet(b *testing.B) {
	cl := NewRedisClient()
	var err error
	for n := 0; n < b.N; n++ {
		err = cl.Set("name", "redis_value_"+fmt.Sprintf("%d", n), 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkRedisGet(b *testing.B) {
	cl := NewRedisClient()
	var err error
	for n := 0; n < b.N; n++ {
		err = cl.Get("name").Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkEtcdSet(b *testing.B) {
	ctx := context.Background()

	cl, err := NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if er := cl.Close(); er != nil {
			log.Fatal(er)
		}
	}()

	for n := 0; n < b.N; n++ {
		_, err = cl.Put(ctx, "name", "etcd_value_"+fmt.Sprintf("%d", n))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkEtcdGet(b *testing.B) {
	ctx := context.Background()

	cl, err := NewEtcdClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if er := cl.Close(); er != nil {
			log.Fatal(er)
		}
	}()

	for n := 0; n < b.N; n++ {
		_, err = cl.Get(ctx, "name")
		if err != nil {
			log.Fatal(err)
		}
	}
}
