## ok模式

使用场景: 在一个表达式返回2个参数的时候使用，第一个参数是一个值或者`nil`，第二个参数是`true/false`或者一个错误`error`

在一个需要赋值的if条件语句中，使用这种模式去检测第二个参数值会让代码显得优雅简洁。

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//声明一个空结构体
	type cat struct {
		Name string
		//带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	//创建cat的实例
	ins := cat{Name: "aaa", Type: 1}
	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	//通过字段名，找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		//从tag中取出需要的tag
		fmt.Println(catType.Tag.get("json"), catType.Tag.get("id"))
	}
}
```
*output*
```go
type 100
```