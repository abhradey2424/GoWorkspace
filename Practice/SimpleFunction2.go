package main

import "fmt"

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func main() {
	s1 := "Hello"
	s2 := "World"
	fmt.Println(s1, s2)
	s1, s2 = swap(s1, s2)
	fmt.Println(s1, s2)
}
