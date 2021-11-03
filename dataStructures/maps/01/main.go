package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)

	myGreeting2 := map[string]string{
		"first":  "Good Morning",
		"second": "Good Afternoon",
	}
	myGreeting["third"] = "Good Evening"

	//delete a key value pair from a map
	delete(myGreeting2, "first")

	fmt.Println(myGreeting, myGreeting2)

	_, ok := myGreeting["second"]
	fmt.Println("exists?:", ok)

	if value, exists := myGreeting2["second"]; exists {
		fmt.Println("value:", value, "exists?", exists)
	} else {
		fmt.Println("this key does not exist on this map")
	}

}
