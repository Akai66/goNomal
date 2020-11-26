package main

import "fmt"

//类型的值方法和指针方法

type Cat struct {
	name           string
	scientificName string
	category       string
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category:%s,name:%s)", cat.scientificName, cat.category, cat.name)
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	fmt.Printf("The cat:%s\n", cat)
	cat.SetName("monster") //指针方法会影响原值，此处会自动取地址,(&cat).SetName("monster")
	fmt.Printf("The cat:%s\n", cat)
	cat.SetNameCopy("little pig") //值方法不会影响原值
	fmt.Printf("The cat:%s\n", cat)

	fmt.Println((&cat).Name())

	//指针类型的方法包括该类型的所有值方法和指针方法
	//基本类型的方法仅包含该类型的所有值方法
	//所以会存在一个指针类型实现了某某接口类型，但是其对应的基本类型却不一定能够作为该接口的实现类型
	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet:%v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet:%v\n", ok)
}
