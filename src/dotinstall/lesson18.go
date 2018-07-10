package main

import (
	"fmt"
)

type user struct{
	name  string
	score int
}


func main(){
	u := new(user)
	
	u.name = "taguchi"
	u.score = 200
	fmt.Println(u)
	
	u2 := user{"fkoji", 100}
	fmt.Println(u2)
}
