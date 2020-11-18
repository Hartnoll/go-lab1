package main

import ( 
	"fmt"
	
)

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) []int {
	mappedSlice := []int{}
	
	for _, a := range slice {
		mappedSlice = append(mappedSlice, f(a))
	}
	return mappedSlice
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	var mappedArray [3]int
	for i := 0; i < 3; i++ {
		mappedArray[i] = (f(array[i]))
	}
	return mappedArray
}

func main() {
	intsArray := [3]int{1,2,3}
	fmt.Println(mapArray(addOne, intsArray))
}
