package main

import "fmt"

func main() {
	mainSlice := []string{"Jan", "Feb", "Mar", "Apr"}
	secSlice := []string{"May", "Jun", "Jul", "Aug"}

	//Appending a Slice to another slice
	mainSlice = append(mainSlice, secSlice...)
	fmt.Println(mainSlice)

	//deleting from a slice - deleting the middle items
	secSlice = append(secSlice[:1], secSlice[3:]...)
	fmt.Println(secSlice)

}
