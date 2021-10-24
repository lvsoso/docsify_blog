### 编译
```shell
go build -ldflags="-s -w" -o server main.go
go build -gcflags=-m -o server main.go
```

### 条件编译

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

### 编译体积

[upx压缩](https://github.com/upx/upx/releases)
```shell
upx -9 server
```



### 测试
```shell
go test -v ./ -test.run  
```

### go mod

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

### cgo

```shell
CGO_CFLAGS="-I/include" CGO_LDFLAGS="-L/lib -lcrypto -lssl" go build  \
 -mod vendor -buildmode=plugin -o=plugin.so x1.go x2.go x3.go x4.go
```

```shell
LD_LIBRARY_PATH="/lib" go test -v ./  -test.run  TestXXX
```

### go build -race
数据竞检测。  

### go build -X

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

### 编译器指示

1. 以// line或/ * line开头的行指示，机器生成的代码中出现。

```go
//line :line
//line :line:col
//line filename:line
//line filename:line:col
/*line :line*/
/*line :line:col*/
/*line filename:line*/
/*line filename:line:col*/
```

2. 以//go:name形式的指示，跟着一段声明作用的代码。

```go
// 不允许逃逸
//go:noescape


// 使用importpath.name作为源码中声明为localname的变量或函数的目标文件符号名称
// 即在别的包中“importpath.name”实际调用定义包中的“localname”，需要引入“unsafe”包。
// 
//go:linkname localname importpath.name
 

// 函数必须在系统栈上运行。
//go:systemstack

// 适用于类型声明，表示一个类型必须不能分配到GC堆上。
//go:noinheap

// 不进行内联
//go:noinline


// 跳过栈溢出检测
//go:nosplit

// 不进行竞态检测
//go:norace
```


### go闭包

> 函数+引用环境=闭包
> 
> 当闭包内没有对外部变量造成修改时，Go 编译器会将自由变量的引用传递优化为直接值传递，避免变量逃逸。
> 
> 在Go中，闭包在底层是一个结构体对象，它包含了函数指针与自由变量。


```shell
// 查看汇编 
go tool compile -S -N -l main.go 

// 逃逸分析
go build -gcflags '-m -m -l' main.go
```
[https://mp.weixin.qq.com/s/kfNcrLZlb5LRi6ILsohwUA](https://mp.weixin.qq.com/s/kfNcrLZlb5LRi6ILsohwUA)


### go 接口
interface 分“runtime.eface”和“runtime.iface”。

```go
// 表示不包含任何方法的空接口
type eface struct {
    _type *_type
    data  unsafe.Pointer
}

// 表示包含方法的接口
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```

接口是否为nil判断
```go
func IsNil(i interface{}) bool {
    vi := reflect.ValueOf(i)
    if vi.Kind() == reflect.Ptr {
        return vi.IsNil()
    }
    return false
}

```

[https://mp.weixin.qq.com/s/vNACbdSDxC9S0LOAr7ngLQ](https://mp.weixin.qq.com/s/vNACbdSDxC9S0LOAr7ngLQ)

[https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/)

[https://halfrost.com/go_interface/](https://halfrost.com/go_interface/)

### new 和 make

make： slice，map，channel，创建后返回值本身，创建slice会带有零值；
new：类型内存分配和初始化，返回指针，不会初始化内部数据结构；


[https://mp.weixin.qq.com/s/tZg3zmESlLmefAWdTR96Tg](https://mp.weixin.qq.com/s/tZg3zmESlLmefAWdTR96Tg)

### fmt
> go/src/fmt
#### 基本符号
```text
// 一般动词
%v 以默认格式输出值,v 代表 Value,当打印结构体 （struct） 时，加号 (%+v) 会添加字段名
%#v 输出值的 Go 语法表示
%T 输出值类型的 Go 语法表示 （T 代表 Type）
%% 输出一个百分号 (%)；不消耗任何值 （因为 % 用作了动词开头，为了区分，输出 % 需要转义）

// 布尔值
%t 单词 true 或者 false （t 代表 True value，真值）

// 整型
%b 以 2 为基数输出 （b 代表 Binary，二进制）
%c 输出对应 Unicode 码点所代表的字符 （c 代表 Character，字符）
%d 以 10 为基数输出 （d 代表 Decimal，十进制）
%o 以 8 为基数输出 （o 代表 Octal，八进制）
%O 以 8 为基数输出，以 0o 为前缀 （同上，大写表示增加前缀）
%q 一个单引号字符，按 Go 语法安全转义。 （q 代表 quote，引号）
%x 以 16 为基数输出，a-f 为小写字母 （x 代表 heXadecimal）
%X 以 16 为基数输出，A-F 为大写字母 （同上，大写表示字母大写）
%U Unicode 格式：如 U+1234；与 "U+%04X" 效果一样 （U 代表 Unicode）

// 浮点数
%b 无小数科学记数法，指数以 2 为底数（但整数部分和指数部分均为十进制数），
 相当于以 strconv.FormatFloat 函数的 'b' 格式，
 如：-123456p-78 （分隔的 p 代表 power (of 2)， 2 的幂）
%e 科学记数法，如：-1.234456e+78 （e 代表 Exponent，指数）
%E 科学记数法，如：-1.234456E+78
%f 无指数的小数，如：123.456 （f 代表 floating-point，浮点数）
%F %f 的同义词
%g 指数较大时等同于 %e ，否则为 %f 。精度在下面讨论。
 （换言之，%g 取 %e 和 %f 中较短的格式表示）
%G 指数较大时等同于 %E ，否则为 %F
%x 十六进制记数法(指数为十进制，底数为 2 )，如：-0x1.23abcp+20
 （与 %b 的区别是，左边的实数为十六进制，而且可以有小数）
%X 十六进制符号大写，如：-0X1.23ABCP+20

// 字符串和字节切片
%s 字符串或字节切片未经解释的字节（uninterpreted bytes） （s 代表 String，字符串），不会对字符序列的内容进行转义
%q 一个双引号字符串，按 Go 语法安全转义
%x 以十六进制数输出，小写，每个字节对应两个字符
%X 以十六进制数输出，大写，每个字节对应两个字符

// 切片
%p 以十六进制数表示的第 1 个元素（下标 0）的地址，以 0x 开头
 （p 代表 Pointer，指针，也就是以指针形式输出地址）

// 指针
%p 十六进制数地址，以 0x 开头
 %b, %d, %o, %x 和 %X 动词也可以用于指针，
 实际上就是把指针的值当作整型数一样格式化。

// %v 的默认格式
bool:      %t
int, int8 等有符号整数:  %d
uint, uint8 等无符号整数: %d, 如果以 %#v 输出则是 %#x
float32, complex64 等:  %g
string:      %s
chan:      %p
指针:      %p

// 复合对象
struct:    {field0 field1 ...}
array, slice:  [elem0 elem1 ...]
maps:    map[key1:value1 key2:value2 ...]
上述类型的指针:  &{}, &[], &map[]

// 宽度与精度
%f  默认宽度，默认精度
%9f  宽度 9，默认精度
%.2f 默认宽度，精度 2
%9.2f 宽度 9, 精度 2
%9.f 宽度 9, 精度 0

// 其他
+ 始终输出数字值的符号（正负号）；
 对 %q(%+q) 保证只输出 ASCII 码 （ASCII 码以外的内容转义）
- 空格填充在右边，而不是左边
# 备选格式：二进制(%#b)加前导 0b ，八进制(%#o)加前导 0 ；
 十六进制(%#x 或 %#X)加前导 0x 或 0X ；%p (%#p) 取消前导 0x ；
 对于 %q ，如果 strconv.CanBackquote 返回true，则输出一个原始（反引号）字符串；
 总是输出 %e, %E, %f, %F, %g 和 %G 的小数点；
 不删除 %g 和 %G 的尾部的零；
 对于 %U (%#U)，如果该字符可打印（printable，即可见字符），则在 Unicode 码后面输出字符，
 例如 U+0078 'x'。
' ' (空格) 为数字中的省略的正号留一个空格 (%d)；
 以十六进制输出字符串或切片时，在字节之间插入空格 (%x, %X)
0 用前导零而不是空格来填充；
 对于数字来说，这会将填充位置移到符号后面
```

#### 规则
```text
除了使用动词 %T 和 %p 输出时，对于实现特定接口的操作数，需要考虑特殊格式化。以下规则按应用顺序排列：

1. 如果操作数是 reflect.Value，则操作数被它所持有的具体值所代替，然后继续按下一条规则输出。

2. 如果操作数实现了 Formatter 接口，则会被调用。在这种情况下，动词和标志的解释由该实现控制。

3. 如果 %v 动词与 # 标志 (%#v) 一起使用，并且操作数实现了 GoStringer 接口，则该接口将被调用。

如果格式 （注意 Println 等函数隐含 %v 动词）对字符串有效 (%s, %q, %v, %x, %X)，则适用以下两条规则：

4. 如果操作数实现了 error 接口，将调用 Error 方法将对象转换为字符串，然后按照动词（如果有的话）的要求进行格式化。
5. 如果操作数实现了 String() string 方法，
```

#### 显示参数索引
```text
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```