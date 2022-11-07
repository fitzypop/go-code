package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var println = fmt.Println

func main() {
	println("Hello Go!")
	println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		println("Hello", name)
	} else {
		log.Fatal(err)
	}
}
