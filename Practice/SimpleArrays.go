package main

import "fmt"

func createArray() {
	var arr [3]int
	arr1 := [3]int{4, 5, 6}
	index := 0
	for index < 3 {
		arr[index] = index
		index = index + 1
	}
	fmt.Println(arr, arr1)
}

func main() {
	createArray()
}
