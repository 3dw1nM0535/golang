package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person object
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person
	reader := strings.NewReader(`{"First": "James", "Last": "Bond", "Age": 10}`)
	json.NewDecoder(reader).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
