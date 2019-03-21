package main

import "fmt"

func main() {
	fmt.Println(isSecondPower(32))
}

func isSecondPower(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n != 2 {
		if n%2 == 0 {
			n /= 2
		} else {
			return false
		}
	}
	return true
}

func isSecondPower2(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
