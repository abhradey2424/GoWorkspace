package main

import "fmt"

func sliceIt() []int {
	var arr [10]int
	index := 0
	for index < 10 {
		arr[index] = index
		index = index + 1
	}
	fmt.Println(arr)
	arr1 := arr[2:8]

	return arr1
}

func main() {
	fmt.Println(sliceIt())
}
