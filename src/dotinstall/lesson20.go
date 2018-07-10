package main

import (
	"fmt"
)

type greeter interface {
	greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet(){
	fmt.Println("konnnichiwa")
}

func (a american) greet(){
	fmt.Println("hello")
}

func main(){
	greeters := []greeter{new(japanese), new(american)}
	
	for _, greeter := range greeters{
		greeter.greet()
	}
	
}
