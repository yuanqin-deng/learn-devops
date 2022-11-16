## :=的详细用法
在Go语言中`:=`运算符被用来声明变量以及给变量赋值，它可以自动推断变量类型；

就像这样：
```go
mystr := "hello world"
```
它等同于：
```go
var mystr string
mystr = "hello world"
```
但是，有些情况你需要知道；

看下面这段代码：
```go
package main

func main(){
	mystr := "hello world"
	println(mystr)
	mystr := "I was just here"
	println(mystr)
}
```

这段代码编译器会报错；因为 := 已经声明变量，相同变量不能在相同的作用域下声明两次，编译器会有这种错误：

```shell
no new variables on left side of :=
```

它告诉你 `:=` 左侧不是新的变量，在这种情况，你应该使用 `=`，因为在前面该变量已经被声明；

可是，在多变量同时赋值时，情况又变得不一样了，看下面这段代码：
```go
package main

func main(){
	mystr := "what is your name?"
	println(mystr)
	mystr,name := "my name is","Robot 1"
	println(mystr,name)
}
```
这段代码能正确编译并运行；

因为在多变量赋值时，`:=` 运算符左侧，只要其中有一个变量是新的，就可以使用 `:=`，这里 `mystr` 使用了两次 `:=` ，因为第二次使用时是在多变量赋值的情况下，且其中包含一个新的变量；

但你不能改变 `mystr` 的类型，因为在第一次使用 `:=` 时，它已经被声明为字符串类型。

如果你使用了 `_` 来忽略其中一个值，那么你要知道，`_` 不属于新的变量。

还有一点，`:=` 不能被用在函数体外，也就是说，当你打算声明一个全局变量时，你只能使用`var`。

