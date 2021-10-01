package main

import (
	"fmt"
	"learningGo/testing/02/quote"
	"learningGo/testing/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	_, m := word.UseCount(quote.SunAlso)
	fmt.Println(m)

	//for k, v := range word.UseCount(quote.SunAlso) {
	//	fmt.Println(v, k)
	//}
}
