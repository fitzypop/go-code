package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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

	// screw styling, I like snake casing
	var v_name string = "Joe"
	// or
	v_last_name := "Fitz"

	println("Hello", v_name, v_last_name)
	var gorilla rune = '🦍'
	println(reflect.TypeOf(gorilla))
}
