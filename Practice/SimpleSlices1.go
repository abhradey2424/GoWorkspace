package main

import "fmt"

func makeSlice() []int {
	arr := make([]int, 10)
	fmt.Println("Initial slice : ", arr)
	index := 0
	for index < 10 {
		arr[index] = index
		index = index + 1
	}
	return arr
}

func main() {
	fmt.Println(makeSlice())
}
