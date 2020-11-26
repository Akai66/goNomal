package main

import "fmt"

//将结构体的普通类型和指针类型分别赋值给接口变量，修改原值的数据，是否会影响接口变量

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name     string
	category string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Category() string {
	return d.category
}

func main() {
	//结构体是值类型，赋值只是复制值的一个副本，所以修改原值不会影响被赋值的接口变量
	dog := Dog{"little pig", "dog"}
	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("The dog name is %q\n", dog.name)
	fmt.Printf("The pet name is %q,category is %q\n", pet.Name(), pet.Category())

	//同上，复制的是值的副本
	dog1 := Dog{"little pig", "dog"}
	dog2 := dog1
	dog1.SetName("monster")
	fmt.Printf("The dog1 name is %q\n", dog1.Name())
	fmt.Printf("The dog2 name is %q\n", dog2.Name())

	//将结构体变量的指针赋值给接口变量，复制得到指针的一个副本，二者指向同一块内存，所以修改原值是会影响被赋值的接口变量
	dog = Dog{"little pig", "dog"}
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog name is %q\n", dog.Name())

	fmt.Printf("The pet name is %q,The pet type is %T\n", pet.Name(), pet)

}
