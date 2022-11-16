## range用法

`range`在`go`中主要是用来做迭代用的，它可以迭代：`array，slice，string，map，channel`
```go
package main
import "fmt"
func main(){
    //数组的遍历
    a := [3]int {1, 2, 3}
    for i, n := range a{
        fmt.Println(i, n)
    }
    //切片的遍历
    b := []int{2, 3, 4}
    for i, n := range b{
        fmt.Println(i, n)
    }
    //map的遍历
    c := map[string]int{"Hello":1, "World":2}
    for k, v := range c{
        fmt.Println(k, v)
    }
}
```
*output*
```go
0 1
1 2
2 3
0 2
1 3
2 4
Hello 1
World 2
```

## 注意事项

1. range会复制对象，而不是直接在原对象上操作。
```go
package main
import "fmt"
func main(){
    a := [3]int {1, 2, 3}
    for _, v := range a{ //复制一份a遍历[1, 2, 3]
        v += 100 //v是复制对象中的值，不会改变a数组元素的值
    }
    fmt.Println(a) //1 2 3
}
```
```go
package main
import "fmt"
func main(){
    a := [3]int {1, 2, 3}
    for i, v := range a{ //i,v从a复制的对象里提取出
        if i == 0{
            a[1], a[2] = 200, 300
            fmt.Println(a) //输出[1 200 300]
        }
        a[i] = v + 100 //v是复制对象里的元素[1, 2, 3]
    }
    fmt.Println(a)  //输出[101, 102, 103]
}
```

2. 使用range迭代遍历引用类型时，底层的数据不会被复制
```go
package main
import "fmt"
func main(){
    a := []int {1, 2, 3} //改成slice
    for i, v := range a{ 
        if i == 0{
            a[1], a[2] = 200, 300
            fmt.Println(a) //[1 200 300]
        }
        a[i] = v + 100 
    }
    fmt.Println(a)  
}
```
*output*
```go
[1 200 300]
[101 300 400]
```

因为切片的内部结构为`struct slice{*point, len, cap}`。

数据部分是一个指针，指向地址，复制对象的时候只是把指针的值复制了，
而不是重新拷贝一块新的内存再把值放进去，所以修改的时候还是修改的原来的值，和C++里的浅拷贝一样。