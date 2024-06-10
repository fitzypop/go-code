package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func HELLO_GO() {
	fmt.Println("Hello Go!")
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		println("Hello", name)
	} else {
		log.Fatal(err)
	}

	// screw styling, I like snake casing
	var v_name string = "Joe"
	// or
	v_last_name := "Fitz"

	println("Hello", v_name, v_last_name)
	gorilla := '🦍'
	println(reflect.TypeOf(gorilla))
}

func main() {
	HELLO_GO()
}
