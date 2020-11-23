//定义方法计算点与点，线段的距离
package geometry

import (
	"math"
    "image/color"
)

type Point struct {
	X, Y float64
}

type ColorPoint struct {
    Point              //嵌入Point结构体
    Color color.RGBA
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(n float64) {
    p.X *= n
    p.Y *= n
}

func (path Path) Distance() float64 {
	var sum float64
	for i := 0; i < len(path); i++ {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
