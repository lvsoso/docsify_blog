## 遇到的问题收集

1. golang 空结构体参数传递，指针地址没有发生变化。
   golang语言参数传递是传值的。对于空结构体，总是`zerobase`地址；
```go
// runtime.newobject(SB)
// runtime.malloc.go
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
 if gcphase == _GCmarktermination {
  throw("mallocgc called with gcphase == _GCmarktermination")
 }

 if size == 0 {
  return unsafe.Pointer(&zerobase)
 }
    ..........
}
``` 

2. golang 中，空结构体作为结构体的字段，在前面和中间不占用空间，在最后则会进行特殊填充，对齐到前一个字段大小，同地址偏移对齐规则。