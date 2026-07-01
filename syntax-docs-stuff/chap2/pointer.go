package main

import (
	"fmt"	
)

func main() {
	v := 3
	incr(&v)
	fmt.Println(v)
}

func incr(v *int) int {
	*v++
	return *v
}
