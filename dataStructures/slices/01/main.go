package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, i)
		//printing the len, and how the capacity grows as the length increases
		fmt.Println("len:", len(mySlice), "cap:", cap(mySlice), "val:", mySlice[i])
	}
}
