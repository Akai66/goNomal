package main

import "fmt"

//给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
//找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。

func numOfBoomerangs(points [][]int) int {
	resNum := 0
	for i, point1 := range points {
		record := map[int]int{}
		for j, point2 := range points {
			if i != j {
				dic := (point1[0]-point2[0])*(point1[0]-point2[0]) + (point1[1]-point2[1])*(point1[1]-point2[1])
				record[dic] += 1
			}
		}
		for _, n := range record {
			resNum += n * (n - 1)
		}
	}
	return resNum
}

func main() {
	points := [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}, {0, 1}, {0, 0}}
	fmt.Println(numOfBoomerangs(points))
}
