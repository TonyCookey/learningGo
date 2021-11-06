package main

import (
	"fmt"
	"sort"
)

//printing from a map in an ordered manner using a different slice
func main() {
	m := map[int]string{
		1: "Tony",
		2: "Ini",
		3: "Ire",
		4: "Kunmi",
		5: "Timilehin",
		6: "Tunde",
	}
	var keys []int
	for k := range m {
		fmt.Println(k)
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
