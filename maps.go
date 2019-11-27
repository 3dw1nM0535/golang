package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour!",
		2: "Buenos dias",
		3: "Bongiorno",
	}

	fmt.Println(myGreeting)
	// delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("Value: ", val)
		fmt.Println("Exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist!")
		fmt.Println("Value: ", val)
		fmt.Println("Exists: ", exists)
	}

	fmt.Println(myGreeting)
}
