减少体积

编译
```shell
go build -ldflags="-s -w" -o server main.go
go build -gcflags=-m -o server main.go
```

条件编译

debug
```go
// +build debug
```

```shell
go build -tags debug -o debug .
```

release
```go
// +build !debug
```

os版本
```go
// +build linux darwin
// +build 386
```

[upx压缩](https://github.com/upx/upx/releases)
```shell
upx -9 server
```



测试
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

```shell
CGO_CFLAGS="-I/include" CGO_LDFLAGS="-L/lib -lcrypto -lssl" go build  \
 -mod vendor -buildmode=plugin -o=plugin.so x1.go x2.go x3.go x4.go
```

```shell
LD_LIBRARY_PATH="/lib" go test -v ./  -test.run  TestXXX
```

```shell
#!/bin/sh

BuildVersion=`git rev-parse --abbrev-ref HEAD`
BuildDate=$(date "+%Y-%m-%d-%H:%M:%S")
CommitHash=`git rev-parse --short HEAD`

TARGET=${BUILD_TARGET}

echo "BuildVersion $BuildVersion"
echo "BuildDate $BuildDate"
echo "CommitHash $CommitHash"

if [ $RELEASE -eq 1 ]; then
  echo "release"
  go build -mod=mod -o ${BUILD_TARGET} -a -ldflags " -extldflags -Wunused-function "-static" -X project/cmd.BuildVersion=$BuildVersion -X project/cmd.BuildDate=$BuildDate -X project/cmd.CommitHash=$CommitHash" main.go
else
  echo "normal"
  go build -mod=mod -o ${BUILD_TARGET} -a -ldflags " -extldflags -Wunused-function -X project/cmd.BuildVersion=$BuildVersion -X project/cmd.BuildDate=$BuildDate -X project/cmd.CommitHash=$CommitHash" main.go
fi

chmod +x ./${BUILD_TARGET}
```

