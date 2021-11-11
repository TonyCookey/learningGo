package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	p1 := people{"tony", "ini", "kunmi", "ire"}
	fmt.Println(p1)

	sort.Strings(p1)

	fmt.Println(p1)
}
