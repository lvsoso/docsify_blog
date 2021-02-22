## Rust

https://github.com/mozilla/sccache

cargo 加速和缓存编译，`vim ~/.cargo/config`

```shell
[source.crates-io]
registry = "https://github.com/rust-lang/crates.io-index"
replace-with = 'ustc'
[source.ustc]
registry = "git://mirrors.ustc.edu.cn/crates.io-index"


[build]
rustc-wrapper = "/home/lv/.cargo/sccache"
```
### 泛型

#### 结构体泛型
```rust
// ############################ 1
// `<T>` 在第一次使用 `T` 前出现
struct SingleGen<T>(T);
struct SGen<T>(T); // 泛型类型 `SGen`。

fn main(){
    {
        // 类型参数显式地指定。
        let _char: SingleGen<char> = SingleGen('a');
        // 类型参数隐式地指定。
        let _t    = SingleGen('a'); 
    }
}

```
#### 函数泛型
```rust
// ############################ 1
// `<T>` 在第一次使用 `T` 前出现
struct SingleGen<T>(T);
struct SGen<T>(T); // 泛型类型 `SGen`。

// ############################ 2
// 函数：在使用类型 T 前给出 <T>，那么 T 就变成了泛型。
// 调用泛型函数有时需要显式地指明类型参量。
fn generic<T>(_s: SGen<T>) {}

fn main(){
    // 为 `generic()` 显式地指定类型参数 `char`。
    generic::<char>(SGen('a'));

    // 为 `generic()` 隐式地指定类型参数 `char`。
    generic(SGen('c'));
}

```
#### impl泛型
```rust

// impl 块实现泛型
struct GenericVal<T>{
     gen_val: T
}; // 泛型类型 `GenericVal`

// `<T>` 必须在类型之前写出来，以使类型 `T` 代表泛型。
impl <T> GenericVal<T> {
    fn value(&self) -> &T { &self.gen_val }
}
```

#### trait泛型
```rust
// 不可复制的类型。
struct Empty;
struct Null;

// `T` 的泛型 trait。
trait DoubleDrop<T> {
    // 定义一个调用者的方法，接受一个额外的参数 `T`，但不对它做任何事。
    fn double_drop(self, _: T);
}

// 对泛型的调用者类型 `U` 和任何泛型类型 `T` 实现 `DoubleDrop<T>` 。
impl<T, U> DoubleDrop<T> for U {
    // 此方法获得两个传入参数的所有权，并释放它们。
    fn double_drop(self, _: T) {}
}

fn main() {
    let empty = Empty;
    let null  = Null;

    // 释放 `empty` 和 `null`。
    empty.double_drop(null);
}
```
#### 约束
约束把泛型类型限制为符合约束的类型。
```rustc
// 定义一个函数 `printer`，接受一个类型为泛型 `T` 的参数，
// 其中 `T` 必须实现 `Display` trait。
fn printer<T: Display>(t: T) {
    println!("{}", t);
}

trait HasArea {
    fn area(&self) -> f64;
}

impl HasArea for Rectangle {
    fn area(&self) -> f64 { self.length * self.height }
}

#[derive(Debug)]
struct Rectangle { length: f64, height: f64 }

// `T` 必须实现 `HasArea`。任意符合该约束的泛型的实例
// 都可访问 `HasArea` 的 `area` 函数
fn area<T: HasArea>(t: &T) -> f64 { t.area() }

```

#### 空约束
即使一个 trait 不包含任何功能，仍然可以用它 作为约束。

```rust
// 标准库trait：Eq 和 Ord。

struct Cardinal;

trait Red {}

impl Red for Cardinal {}

// 这些函数只对实现了相应的 trait 的类型有效。
// 事实上这些 trait 内部是空的，但这没有关系。
fn red<T: Red>(_: &T)   -> &'static str { "red" }

fn main() {
    let cardinal = Cardinal;

    println!("A cardinal is {}", red(&cardinal));
}
```

### 多重约束
多重约束用 + 连接;类型之间使用 , 隔开。
```rust
use std::fmt::{Debug, Display};

fn compare_prints<T: Debug + Display>(t: &T) {
    println!("Debug: `{:?}`", t);
    println!("Display: `{}`", t);
}

fn compare_types<T: Debug, U: Debug>(t: &T, u: &U) {
    println!("t: `{:?}", t);
    println!("u: `{:?}", u);
}
```

#### where 分句
```rust
impl <A: TraitB + TraitC, D: TraitE + TraitF> MyTrait<A, D> for YourType {}

// 使用 `where` 从句来表达约束
impl <A, D> MyTrait<A, D> for YourType where
    A: TraitB + TraitC,
    D: TraitE + TraitF {}
```

#### 关联项
“关联项”指与多种类型的项有关的一组规则。它是 trait 泛型的扩展，允许在 trait 内部定义新的项。
> 一种将类型占位符与 trait 联系起来的做法，这样 trait 中的方法签名中就可以使用这些占位符类型。trait 的实现会指定在该实现中那些占位符对应什么具体类型。

```rust
struct Container(i32, i32);

// 这个 trait 检查给定的 2 个项是否储存于容器中
// 并且能够获得容器的第一个或最后一个值。
trait Contains<A, B> {
    fn contains(&self, _: &A, _: &B) -> bool; // 显式地要求 `A` 和 `B`
    fn first(&self) -> i32; // 未显式地要求 `A` 或 `B`
    fn last(&self) -> i32;  // 未显式地要求 `A` 或 `B`
}

impl Contains<i32, i32> for Container {
    // 如果存储的数字和给定的相等则为真。
    fn contains(&self, number_1: &i32, number_2: &i32) -> bool {
        (&self.0 == number_1) && (&self.1 == number_2)
    }
    // 得到第一个数字。
    fn first(&self) -> i32 { self.0 }
    // 得到最后一个数字。
    fn last(&self) -> i32 { self.1 }
}

// 容器 `C` 就包含了 `A` 和 `B` 类型。
fn difference<A, B, C>(container: &C) -> i32 where
    C: Contains<A, B> {
    container.last() - container.first()
}
```

```rust
struct Container(i32, i32);

// 这个 trait 检查给定的 2 个项是否储存于容器中
// 并且能够获得容器的第一个或最后一个值。
trait Contains {
    // 在这里定义可以被方法使用的泛型类型。
    type A;
    type B;

    fn contains(&self, _: &Self::A, _: &Self::B) -> bool;
    fn first(&self) -> i32;
    fn last(&self) -> i32;
}

impl Contains for Container {
    // 指出 `A` 和 `B` 是什么类型。如果 `input`（输入）类型
    // 为 `Container(i32, i32)`，那么 `output`（输出）类型
    // 会被确定为 `i32` 和 `i32`。
    type A = i32;
    type B = i32;

    // `&Self::A` 和 `&Self::B` 在这里也是合法的类型。
    fn contains(&self, number_1: &i32, number_2: &i32) -> bool {
        (&self.0 == number_1) && (&self.1 == number_2)
    }

    // 得到第一个数字。
    fn first(&self) -> i32 { self.0 }

    // 得到最后一个数字。
    fn last(&self) -> i32 { self.1 }
}

fn difference<C: Contains>(container: &C) -> i32 {
    container.last() - container.first()
}

fn main() {
    let number_1 = 3;
    let number_2 = 10;

    let container = Container(number_1, number_2);

    println!("Does container contain {} and {}: {}",
        &number_1, &number_2,
        container.contains(&number_1, &number_2));
    println!("First number: {}", container.first());
    println!("Last number: {}", container.last());

    println!("The difference is: {}", difference(&container));
}
```


#### 虚类型参数
虚类型（phantom type）参数是一种在运行时不出现，而在（且仅在）编译时进行静态检查的类型参数。
```rust

```


### rustc

```shell
rustc --crate-type=lib yyy.rs
rustc xxx.rs --extern yyy=./libyyy.rlib
```

### Cargo

Cargo toml

```toml
[package]
name = "hello_cargo"
version = "0.1.0"
authors = ["wuqz <wuqz@siccs.cn>"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
clap = "2.27.1" # 来自 crates.io
rand = { git = "https://github.com/rust-lang-nursery/rand" } # 来自网上的仓库
bar = { path = "../bar" } # 来自本地文件系统的路径

```

### attribute
 属性语法
```rust
#![crate_attribute]

#[item_attribute]

#[attribute = "value"]
#[attribute(key = "value")]
#[attribute(value)]

#[attribute(value, value2)]

#[attribute(value, value2, value3,
    value4, value5)]
```

dead_code

```rust
#[allow(dead_code)]
```

file type
```rust
// 这个 crate 是一个库文件
#![crate_type = "lib"]
// 库的名称为 “rary”
#![crate_name = "rary"]
```

条件编译
cfg 属性：在属性位置中使用 #[cfg(...)]
cfg! 宏：在布尔表达式中使用 cfg!(...)

```rust
#[cfg(target_os="linux")]
#[cfg(not(target_os = "linux"))]

// rustc --cfg some_condition
#[cfg(some_condition)]
fn conditional_function() {
    println!("condition met!")
}

fn main(){
    if  cfg!(target_os = "linux"){
        println!("linux");
    }
}

```

```rust
#[derive(Debug)]

```

### Error Handling
panic:主要用于测试，以及处理不可恢复的错误。
Option:Some(T), None值是可选的、或者缺少值并不是错误的情况准备的。
Result:当错误有可能发生，且应当由调用者处理时,可以 unwrap 然后使用 expect。

#### Option
##### map
```rust
fn process(food: Option<Food>) -> Option<Cooked> {
    food.map(|f| Peeled(f))
        .map(|Peeled(f)| Chopped(f))
        .map(|Chopped(f)| Cooked(f))
}
```
##### and_then
 使用被 Option 包裹的值来调用其输入函数并返回结果。 如果 Option 是 None，那么它返回 None。
```rust
fn cookable_v2(food: Food) -> Option<Food> {
    have_ingredients(food).and_then(have_recipe)
}
```
#### Result
> Ok<T>：找到 T 元素。

> Err<E>：找到 E 元素，E 即表示错误的类型。

>  unwrap()：要么拿到T，要么panic。

> ?：try! ，需要 unwrap 并且不产生 panic。

#### 处理多种错误类型(unwrap)
``` rust
fn double_first(vec: Vec<&str>) -> i32 {
    let first = vec.first().unwrap(); // 生成错误 1
    2 * first.parse::<i32>().unwrap() // 生成错误 2
}

fn main() {
    let numbers = vec!["42", "93", "18"];
    let empty = vec![];
    let strings = vec!["tofu", "93", "18"];

    println!("The first doubled is {}", double_first(numbers));

    println!("The first doubled is {}", double_first(empty));
    // 错误1：输入 vector 为空

    println!("The first doubled is {}", double_first(strings));
    // 错误2：此元素不能解析成数字
}
```

#### 定义一个错误类型
- 用同一个类型代表了多种错误
- 向用户提供了清楚的错误信息
- 能够容易地与其他类型比较
- 能够容纳错误的具体信息
- 能够与其他错误很好地整合



### macro

```rust

```

