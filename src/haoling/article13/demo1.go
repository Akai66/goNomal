package main

import "fmt"

type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}

type Animal struct {
	scientificName string
	AnimalCategory //嵌入字段，又称为"匿名字段"
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (a Animal) String() string {
	return fmt.Sprintf("%s (category:%s)", a.scientificName, a.AnimalCategory)
}

//多层嵌套，名称相同的方法会逐层屏蔽
type Cat struct {
	name string
	Animal
}

func (c Cat) String() string {
	return fmt.Sprintf("%s (category:%s,name:%s)\n", c.scientificName, c.AnimalCategory, c.name)
}

func main() {
	//go可以面向对象编程，但是没有继承的概念，只是通过嵌入字段的方式，实现类型之间的组合
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category:%s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal:%s\n", animal)
	fmt.Println(animal.Category())

	cat := Cat{
		name:   "棉花糖",
		Animal: animal,
	}
	fmt.Printf("The animal:%s\n", cat)
}
