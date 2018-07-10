package main

import (
	"fmt"
)

func main(){
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
	
	i := 0
	for{
		i++
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}
