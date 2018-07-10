package main

import (
	"fmt"
)

func main(){
	signal := "red"

	switch signal{
		case "red":
			fmt.Println("stop")
		case "yellow":
			fmt.Println("caution")
		case "blue"
			fmt.Println("go")
		default:
			fmt.Println("wrong signal")
	}
}
