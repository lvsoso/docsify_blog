### coding

#### 代码风格

##### 格式

- gofmt

- goimports



```go
import (
    // go 标准包
    "fmt"
    
    // 第三方包
    "github.com/jinzhu/gorm"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    
    // 匿名包单独分组，并对匿名包引用进行说明
    // import mysql driver
     _ "github.com/jinzhu/gorm/dialects/mysql"
    
    // 内部包
     v1 "github.com/marmotedu/api/apiserver/v1"
     metav1 "github.com/marmotedu/apimachinery/pkg/meta/v1"
    "github.com/marmotedu/iam/pkg/cli/genericclioptions"
 )
```

- 函数外部声明必须使用 var ，不要采用 :=

- 使用 &T{}代替 new(T)。
- 尽可能指定容器容量，以便为容器预先分配内存。
- 在顶层，使用标准 var 关键字。请勿指定类型，除非它与表达式的类型不同。
- 对于未导出的顶层常量和变量，使用 _ 作为前缀。
- 嵌入式类型（例如 mutex）应位于结构体内的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开。

##### 错误处理

- 错误要单独判断，不与其他逻辑组合判断。
- 错误描述建议
  - 告诉用户他们可以做什么，而不是告诉他们不能做什么。
  - 当声明一个需求时，用 must 而不是 should。例如，must be greater than 0、must match regex '[a-z]+'。
  - 当声明一个格式不对时，用 must not。例如，must not contain。
  - 当声明一个动作时用 may not。例如，may not be specified when otherField is empty、only name may be specified。
  - 引用文字字符串值时，请在单引号中指示文字。例如，ust not contain '..'。
  - 当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than request。
  - 指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than or equal to 0 (不要用 larger than、bigger than、more than、higher than)。
  - 指定数字范围时，请尽可能使用包含范围。
  - 建议 Go 1.13 以上，error 生成方式为 fmt.Errorf("module xxx: %w", err)。
  - 错误描述用小写字母开头，结尾不要加标点符号，例如：

##### panic处理

- 在业务逻辑处理中禁止使用 panic。
- 在 main 包中，**只有当程序完全不可运行时使用 panic**，例如无法打开文件、无法连接数据库导致程序无法正常运行。

- 在 main 包中，使用 log.Fatal 来记录错误，这样就可以由 log 来结束程序，或者将 panic 抛出的异常记录到日志文件中，方便排查问题。
- **可导出的接口一定不能有 panic**。

- 包内建议采用 error 而不是 panic 来传递错误。

##### 单元测试

- 每个**重要**的可导出函数都要编写测试用例。
- 如果存在 func (b *Bar) Foo ，单测函数可以为 func TestBar_Foo。

##### 类型断言

``` go
// bad
t := n.(int)

// good
t, ok := n.(int)
if !ok {
// error handling
}

// normal code
```

#### 命名

##### 包名

- 包名必须和目录名一致，尽量采取有意义、简短的包名，不要和标准库冲突。
- **包名全部小写**，没有大写或下划线，使用多级目录来划分层级。

- 项目名可以通过中划线来连接多个单词。
- **包名以及包所在的目录名，不要使用复数**，例如，是net/utl，而不是net/urls。

- 不要用 common、util、shared 或者 lib 这类宽泛的、无意义的包名。
- 包名要简单明了，例如 net、time、log。

##### 文件命名

- 文件名要简短有意义。

##### 结构体

- 避免使用 Data、Info 这类无意义的结构体名。

##### 变量

- 若变量类型为 bool 类型，则名称应以 Has，Is，Can 或 Allow 开头
- 局部变量应当尽可能短小

##### 接口

- 接口命名的规则，基本和结构体命名规则保持一致：
- 单个函数的接口名以 “er"”作为后缀（例如 Reader，Writer），有时候可能导致蹩脚的英文，但是没关系。

- 两个函数的接口名以两个函数名命名，例如 ReadWriter。
- 三个以上函数的接口名，类似于结构体名。

##### 常量

- 常量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。

- 如果是枚举类型的常量，需要先创建相应类型

##### Error

- Error 类型应该写成 FooError 的形式。
- Error 变量写成 ErrFoo 的形式。



#### 注释

- 每个可导出的名字都要有注释，该注释对导出的变量、函数、结构体、接口等进行简要介绍。

- 全部使用单行注释，禁止使用多行注释。
- 和代码的规范一样，单行注释不要过长，禁止超过 120 字符，超过的请使用换行展示，尽量保持格式优雅。
- 注释必须是完整的句子，以需要注释的内容作为开头，句点作为结尾，格式为 // 名称 描述. 。
- **注释掉的代码**在提交 code review 前都**应该被删除**，否则应该说明为什么不删除，并给出后续处理建议。



```go
// PrintFlags logs the flags in the flagset.
funcPrintFlags(flags *pflag.FlagSet) {
// normal code
}

// 多段注释之间可以使用空行分隔加以区分
// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman

// Package 包名 包描述
// Package genericclioptions contains flags which can be added to you command, bound, completed, and produce
// useful helper functions.
package genericclioptions


// 变量名 变量描述
// ErrSigningMethod defines invalid signing method error.
var ErrSigningMethod = errors.New("Invalid signing method")


// Code must start with 1xxxxx. 
const ( 
    
    // ErrSuccess - 200: OK. 
    ErrSuccess int = iota + 100001
    
    // ErrUnknown - 500: Internal server error. 
    ErrUnknown 
    
	// ErrBind - 400: Error occurred while binding the request body to the struct. 
 	ErrBind
    
	// ErrValidation - 400: Validation failed. 
 	ErrValidation
）
    
// 结构体名 结构体描述
// User represents a user restful resource. It is also used as gorm model.
type User struct {
    // Standard object's metadata.
     metav1.ObjectMeta `json:"metadata,omitempty"`
     Nickname string `json:"nickname" gorm:"column:nickname"`
     Password string `json:"password" gorm:"column:password"`
     Email string `json:"email" gorm:"column:email"`
     Phone string `json:"phone" gorm:"column:phone"`
     IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin"`
}


// 函数名 函数描述
// BeforeUpdate run before update database record.
func(p *Policy)BeforeUpdate()(err error) {
// normal code
return nil
}

// 类型名 类型描述
// Code defines an error code type.
type Code int

```



#### 类型

##### 字符串

空字符串判断。

```go
// good
if len(s)== 0 {
// normal code
}
```

[]byte/string 相等比较。

```go
// good
var s1 []byte
var s2 []byte
...
bytes.Compare(s1, s2) == 0
bytes.Compare(s1, s2) != 0
```

字符串是否包含子串或字符。

```go
strings.Index(s, subStr) > -1
strings.IndexAny(s, char) > -1
strings.IndexRune(s, r) > -1
```

去除前后子串。

```go
// good
var s1 = "a string value"
var s2 = "a "

var s3 string
if strings.HasPrefix(s1, s2) {
 s3 = s1[len(s2):]
}

```

复杂字符串使用 raw 字符串避免字符转义。**？？？**

```go
// good
regexp.MustCompile(`\.`)
```



##### 切片

对于空切片，需要判断是否为nil或在长度为0

```go
// good
if slice != nil && len(slice) == 0 {
	// normal code
}
```

##### 结构体

struct 以多行格式初始化。

#### 控制结构

- if 对于 bool 类型的变量，应直接进行真假判断。
- 采用短声明建立局部变量。
- 不要在 for 循环里面使用 defer。
- range 按需使用。
- switch必须要有 default。
- 不要用goto。

#### 函数

- 传入变量和返回变量都以小写字母开头。
- 尽量用值传递，非指针传递。
- 传入参数是 map、slice、chan、interface ，不要传递指针。
- 如果函数返回相同类型的两个或三个参数，或者如果从上下文中不清楚结果的含义，使用命名返回，**其他情况不建议使用命名返回**
  - 参数数量均不能超过 5 个。
  - 多返回值最多返回三个，超过三个请使用 struct。
- 当存在资源创建时，应紧跟 defer 释放资源
- 先判断是否错误，再 defer 释放资源

##### 方法的接收器

- 推荐以类名第一个英文首字母的小写作为接收器的命名。
- 接收器的命名在函数超过 20 行的时候不要用单字符。

- 接收器的命名不能采用 me、this、self 这类易混淆名称。

##### 嵌套

- 嵌套深度不能超过 4 层

##### 变量命名

- 变量声明尽量放在变量第一次使用的前面，遵循就近原则。

- 如果魔法数字出现超过两次，则禁止使用，改用一个常量代替。



#### GOPATH

- Go 1.11 之后，弱化了 GOPATH 规则，已有代码（很多库肯定是在 1.11 之前建立的）肯定符合这个规则，建议保留 GOPATH 规则，便于维护代码。
- 建议只使用一个 GOPATH，不建议使用多个 GOPATH。如果使用多个 GOPATH，编译生效的 bin 目录是在第一个 GOPATH 下。

#### 依赖管理

- 使用 Go Modules 作为依赖管理的项目时，不建议提交 vendor 目录。
- 使用 Go Modules 作为依赖管理的项目时，必须提交 go.sum 文件。

#### 接口验证

```go
type LogHandler struct {
 	h http.Handler
	log *zap.Logger
}
var _ http.Handler = LogHandler{}
```



#### 性能

- 如果没有特殊需要，需要修改时多使用 []byte。
- 优先使用 strconv 而不是 fmt。
- append 要小心自动分配内存，append 返回的可能是新分配的地址。
- 如果要直接修改 map 的 value 值，则 value 只能是指针，否则要覆盖原来的值。
- map 在并发中需要加锁。
- 编译- 过程无法检查 interface{} 的转换，只能在运行时检查，小心引起 panic。

