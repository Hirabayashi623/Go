package main

import (
	"fmt"
)

type user struct{
	name  string
	score int
}

func (u user)show() {
	fmt.Printf("name: %s, score: %d\n", u.name, u.score)
}

func (u *user)add() {
	u.score++
}


func main(){
	u := new(user)
	
	u.name = "taguchi"
	u.score = 200
	u.show()
	u.add()
	u.show()
}
