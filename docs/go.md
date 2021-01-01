```shell
go test -v ./ -test.run  
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

