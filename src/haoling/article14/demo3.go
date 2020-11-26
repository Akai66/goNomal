package main

import (
	"fmt"
)

type Animal interface {
	ScientificName() string
	Category() string
}

type Named interface {
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scientificName string
	category       string
}

func (d Dog) ScientificName() string {
	return d.scientificName
}

func (d Dog) Category() string {
	return d.category
}

func main() {
	petTag := PetTag{name: "little pig"}
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface Named:%v\n", ok)
	dog := Dog{
		petTag,
		"Labrador Retriever",
		"dog",
	}
	fmt.Println(dog.Name())
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal:%v\n", ok)
	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named:%v\n", ok)
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet:%v\n", ok)
}
