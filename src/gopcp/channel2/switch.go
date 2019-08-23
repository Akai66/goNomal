package main

import (
	"fmt"
)

func main() {
	var v interface{}

	v = true

	//类型断言
	//i1,ok := interface{}(v1).(I1)
	//如果断言的结果为true，i1就是经过类型转换后的I1类型的值，如果断言结果为false，会引发一个恐慌，但是采用上述方式，运行时恐慌就不会发生，且i1是I1类型的零值
	//注意：v.(type) 可以获得一个程序实体的具体类型，但是这种用法只能在类型switch语句的表达式中使用
	switch v.(type) {
	case string:
		fmt.Printf("The string is %s\n", v)
	case int, int16, int32:
		fmt.Printf("The integer is %d\n", v)
	default:
		fmt.Printf("unSupported value type : %s \n", v.(type))
	}
}
