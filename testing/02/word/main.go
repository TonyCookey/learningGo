//Package word has functions that compute the number in a given string
package word

import (
	"strings"
)

// UseCount counts how many times each word appears in the string
func UseCount(s string) (map[string]int, int) {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m, len(m)
}

//Count returns the number of words in a string
func Count(s string) int {
	xs := strings.Fields(s)

	return len(xs)
}
