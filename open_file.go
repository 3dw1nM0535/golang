package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln("Program reading file failed")
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Program reading file content failed")
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
