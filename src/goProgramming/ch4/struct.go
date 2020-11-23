//结构体嵌入匿名成员
package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{1, 1}, 5}, 20}
	fmt.Printf("%#v\n", w)
	w.X, w.Y, w.Radius = 2, 2, 6
	fmt.Printf("%#v\n", w)
}
