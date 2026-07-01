package main

import(
	"fmt"
)

type team int8

func main() { 
	var v team = 5
	x := calc(v)
	fmt.Println(x)
}

func calc(t team) team {
	t++
	return t
}

