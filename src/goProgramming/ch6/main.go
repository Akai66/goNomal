package main

import (
	"fmt"
	"goProgramming/ch6/geometry"
    "image/color"
)

func main() {
	path := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
    fmt.Println(path.Distance())
    red := color.RGBA{255,0,0,255}
    blue := color.RGBA{0,0,255,255}
    cp := geometry.ColorPoint{geometry.Point{1,1},red}
    cq := geometry.ColorPoint{geometry.Point{5,4},blue}
    fmt.Println(cp.Distance(cq.Point))
    cp.ScaleBy(2)
    cq.ScaleBy(2)
    fmt.Println(cp.Distance(cq.Point))
}
