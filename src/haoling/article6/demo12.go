package main

import (
	"fmt"
	"os"
)

var container = []string{"a", "b", "c"}

func main() {
	container := map[int64]string{0: "a", 1: "b", 2: "c"}
	//第一种方式,断言
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int64]string)
	if !(ok1 || ok2) {
		fmt.Fprintf(os.Stderr, "unsupport container type:%T\n", container)
	}
	fmt.Printf("ele:%s,container type:%T\n", container[1], container)
	ele, err := getElement(container)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Printf("ele:%s,container type:%T\n", ele, container)
}

func getElement(container interface{}) (ele string, err error) {
	switch t := container.(type) {
	case []string:
		ele = t[1]
	case map[int64]string:
		ele = t[1]
	default:
		err = fmt.Errorf("unsupport container type:%T\n", t)
		return
	}
	return
}
