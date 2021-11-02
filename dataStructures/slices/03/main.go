package main

import "fmt"

func main() {
	studentRecords := make([][]string, 0)

	//appending slices to a slice
	student1 := []string{"Ben", "Foster", "Computer Science"}
	studentRecords = append(studentRecords, student1)
	student2 := []string{"China", "Audrey", "Computer Engineering"}
	studentRecords = append(studentRecords, student2)

	fmt.Println(studentRecords)
}
