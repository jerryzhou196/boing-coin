package main

import (
	"fmt"

	"example.com/greetings"
)

func main(){
	message := greetings.Hello("ballsack")
	fmt.Println(message)
}
