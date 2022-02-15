# redis vs etcd with go


```
git clone https://github.com/romanitalian/redis_vs_etc_with_go
cd redis_vs_etc_with_go
```

```
go test -bench=. -benchmem
go test -bench=Redis -benchmem
go test -bench=Etcd -benchmem


BenchmarkRedisSet-8   	    1526	    802456 ns/op	     249 B/op	       9 allocs/op
BenchmarkRedisGet-8   	    1552	    818410 ns/op	     196 B/op	       7 allocs/op
BenchmarkEtcdSet-8    	      80	  27729002 ns/op	   14010 B/op	     124 allocs/op
BenchmarkEtcdGet-8    	    9590	    125109 ns/op	    7070 B/op	     125 allocs/op
```


```
go run main.go


Redis 'hello world'
	PONG
	redis is awesome

etcd 'hello world'
	res.Kvs[0].Version: 14
	res.Kvs[0].Value: etcd is great

Time execution
	redis set: 7.825733292s
	redis get: 7.625631125s
	etcd set: 4m21.974348792s
	etcd get: 1.302687666s

```

```
goos: darwin
goarch: arm64

system_profiler SPHardwareDataType
Hardware:

    Hardware Overview:
      Model Name: MacBook Pro
      Chip: Apple M1
      Total Number of Cores: 8 (4 performance and 4 efficiency)
      Memory: 16 GB
```
