package main
import (
	"fmt"
    "math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		fmt.Println("参数必须大于0")
	}
	z := 1.0
	for {
		lastZ := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-lastZ) < 0.00000000001 {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
