## make的使用

`golang` 分配内存主要有内置函数`new`和`make`

`make`函数用于初始化slice、channel和map, 如果只用var声明，不用make初始化，变量对应的值为nil.

首先来看下`make`有以下三种不同的用法：

1. make(map[string]string)

第一种用法，即缺少长度的参数，只传类型，这种用法只能用在类型为`map`或`chan`的场景，例如`make([]int)`是会报错的。这样返回的空间长度都是默认为0的。

2. make([]int, 2)

第二种用法，指定了长度，例如`make([]int, 2)`返回的是一个长度为2的`slice`

3. make([]int, 2, 4)

第三种用法，第二参数指定的是切片的长度，第三个参数是用来指定预留的空间长度，例如`a := make([]int, 2, 4)`, 这里值得注意的是返回的切片a的总长度是4，预留的意思并不是另外多出来4的长度，其实是包含了前面2个已经切片的个数的。所以举个例子当你这样用的时候 `a := make([]int, 4, 2)`，就会报语法错误。

因此，当我们为`slice`分配内存的时候，应当尽量预估到`slice`可能的最大长度，通过给`make`传第三个参数的方式来给`slice`预留好内存空间，
这样可以避免二次分配内存带来的开销，大大提高程序的性能。

而事实上，我们其实是很难预估切片的可能的最大长度的，这种情况下，当我们调用`append`为`slice`追加元素时，`golang`为了尽可能的减少二次分配内存，
并不是每一次都只增加一个单位的内存空间，而且遵循这样一种扩容机制：

当有预留的未使用的空间时，直接对未使用的空间进行切片追加，当预留的空间全部使用完毕的时候，扩容的空间将会是当前的`slice`长度的一倍，
例如当前`slice`的长度为4，进行一次`append`操作之后，`cap(a)`返回的长度将会是8.来看下面这段演示代码:

```go
package main

import (
        "fmt"
)

func main() {
        a :=  make([]int, 0)
        n := 20
        for i := 0; i < n; i++ {
                a = append(a, 1)
                fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
        }
}

Output:
len=1 cap=1  // 第一次扩容
len=2 cap=2 // 第二次扩容
len=3 cap=4 // 第三次扩容
len=4 cap=4
len=5 cap=8 // 第四次扩容
len=6 cap=8
len=7 cap=8
len=8 cap=8
len=9 cap=16 // 第五次扩容
len=10 cap=16
len=11 cap=16
len=12 cap=16
len=13 cap=16
len=14 cap=16
len=15 cap=16
len=16 cap=16
len=17 cap=32 // 第六次扩容
len=18 cap=32
len=19 cap=32
len=20 cap=32
```
以上测试结果表明，每次扩容后，内存空间长度会变为原来的两倍。



