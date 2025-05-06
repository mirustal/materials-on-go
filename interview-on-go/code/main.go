package main

import "fmt"

func main() {
	slice := make([]int, 3, 4)
	appendingSlice(slice[:1])
	fmt.Println(slice, len(slice), cap(slice))
}

func appendingSlice(slice []int) {
	fmt.Printf("slice: %v len: %v, cap: %v", slice, len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Printf("\nslice: %v len: %v, cap: %v", slice, len(slice), cap(slice))
}
