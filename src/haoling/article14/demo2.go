package main

import (
	"fmt"
	"reflect"
)

//将nil赋值给接口变量，接口变量并不为nil，接口变量的存储结构为iface，其中有两个指针，一个指向动态类型信息，一个指向动态值
//只有将nil字面值直接赋值给接口变量或者仅定义接口变量(未赋值),这两种情况，接口变量才为nil

type Pet interface {
	Name() string
}

type Dog struct {
	name string
}

func (d Dog) Name() string {
	return d.name
}

func main() {
	var dog1 *Dog
	fmt.Printf("The dog1 is %v\n", dog1)
	dog2 := dog1
	fmt.Printf("The dog2 is %v\n", dog2)

	var pet Pet = dog1
	if pet == nil {
		fmt.Println("The pet is nil")
	} else {
		fmt.Println("The pet is not nil")
	}

	fmt.Printf("The type of pet is %T\n", pet)
	fmt.Printf("The type of pet is %s\n", reflect.TypeOf(pet).String())
	fmt.Printf("The type of dog2 is %T\n", dog2)

	warp := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}

	pet = warp(dog2)

	if pet == nil {
		fmt.Println("The pet is nil")
	} else {
		fmt.Println("The pet is not nil")
	}

}
