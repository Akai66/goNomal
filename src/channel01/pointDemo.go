package main

import "fmt"

func main(){
	var a int = 400
	var p1 *int = &a
	fmt.Printf("指针p1的值为:%v,指向的值为:%d\n",p1,*p1)
	*p1 = 200
	fmt.Printf("a的值为:%d\n",a)
	var b float64 = 3.2
	var p2 *float64 = &b
	fmt.Printf("指针p2的值为:%v,指向的值为:%f\n",p2,*p2)
	*p2 = 2.4
	fmt.Printf("b的值为:%f\n",b)
}