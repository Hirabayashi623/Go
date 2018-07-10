package main

import (
	"fmt"
)

func main(){
	s := []int{2,3,4}
	
	for i, v := range s {
		fmt.Println(i, v)
	}
	
	for _, v := range s {
		fmt.Println(v)
	}
	
	m := map[string]int{"taguchi":200, "fkoji":100}
	
	for k, v := range m {
		fmt.Println(k, v)
	}
}
