package main 

import (
//	"strings"
	"fmt"
	"os"
)

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	// we can also let Println do the formatting for us 
	fmt.Println(os.Args[1:])
}

