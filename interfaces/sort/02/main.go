package main

import (
	"fmt"
	"sort"
)

func main() {
	p1 := []string{"tony", "ini", "kunmi", "ire"}
	fmt.Println(p1)

	sort.Strings(p1)

	fmt.Println(p1)
}
