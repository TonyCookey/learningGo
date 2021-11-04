package main

import "fmt"

func main() {
	ballondor := map[int]string{
		1: "Cristiano Ronaldo",
		2: "Lionel Messi",
		3: "Robert Lewandowski",
		4: "Neymar",
		5: "Kylian Mbappe Lottin",
	}

	for key, value := range ballondor {
		//maps are unordered
		fmt.Println("Ballon D'Or Nominee", key, "---", "value:", value)
	}
}
