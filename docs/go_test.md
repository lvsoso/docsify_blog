### Go Test

[https://zhuanlan.zhihu.com/p/80578541](https://zhuanlan.zhihu.com/p/80578541)

- parallel 并发数量
- count 数量

```shell
go test -bench="." -parallel 50  -count 1000000 
```

- benchtime 时间

```shell
go test -bench=BenchmarkSHA256Parallel  -benchtime=20s
```

#### BenchmarkSHA256Parallel

```go
func BenchmarkSHA256Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		data := make([]byte, 2048)
		rand.Read(data)
		for pb.Next() {
			h := sha256.New()
			h.Write(data)
			h.Sum(nil)
		}
	})
}

```

