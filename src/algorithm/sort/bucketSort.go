package main

import "fmt"

func main() {
	arr := []int{10, 8, 19, 2, 7, 3, 5, 4, 2, 1, 25, 13, 14, 17, 11, 20}
	bucketSort(arr, 5)
	fmt.Println(arr)
}

func bucketSort(arr []int, num int) {
	if len(arr) <= 1 {
		return
	}
	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	gap := (max-min)/num + 1
	buckets := make([][]int, num)
	for _, v := range arr {
		index := (v - min) / gap
		buckets[index] = append(buckets[index], v)
	}
	j := 0
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertSort(bucket)
			copy(arr[j:], bucket)
			j += len(bucket)
		}
	}
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		preIndex := i - 1
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}

}
