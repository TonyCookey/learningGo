package main

import "fmt"

func main() {
	num := []string{"Hello", "World"}
	str := "Hello World"
	xbs := []byte(str)
	emptyInterface(num, str, xbs)
}
func emptyInterface(x ...interface{}) {
	fmt.Println(x)
}
