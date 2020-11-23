//map的顺序是不定的，如何有序遍历map
package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{"Bob": 24, "Jack": 18, "Rose": 30, "Mary": 12}
	//使用切片保存key
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	//对name进行排序
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
