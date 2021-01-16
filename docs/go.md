```shell
go test -v ./ -test.run  
```



go mod

```shell
go mod why -m all  解释为什么需要依赖
go mod graph       打印模块依赖图
go mod verify      校验依赖
```



```text

replace (
	xxx  => "../xxx"
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

