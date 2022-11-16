## hash

> hash包提供hash函数的接口。
### type Hash
```go
type Hash interface {
    // 通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据，永远不返回错误
    io.Writer
    // 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
    Sum(b []byte) []byte
    // 重设hash为无数据输入的状态
    Reset()
    // 返回Sum会返回的切片的长度
    Size() int
    // 返回hash底层的块大小；Write方法可以接受任何大小的数据，
    // 但提供的数据是块大小的倍数时效率更高
    BlockSize() int
}
```
> Hash是一个被所有hash函数实现的公共接口。

### type Hash32
```go
type Hash32 interface {
    Hash
    Sum32() uint32
}
```
> Hash32是一个被所有32位hash函数实现的公共接口。

### type Hash64
```go
type Hash64 interface {
    Hash
    Sum64() uint64
}
```
> Hash64是一个被所有64位hash函数实现的公共接口。

- adler32
- crc32
- crc64
- fnv

