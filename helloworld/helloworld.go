package main

import (
	"fmt"
	"os"
	"strconv"
)

func getOption() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error getOption", err)
		}
	}()

	argsWithProg := os.Args
	if len(argsWithProg) == 1 {
		return 10
	}

	v, err := strconv.Atoi(argsWithProg[1]) //strconv.Aoti(string(argsWithProg[1]))
	if err != nil {
		return 10
	}

	return v
}

// type aaa struct {
// 	c int
// 	b int
// }

func main() {
	var c = getOption()
	fmt.Printf("Hello World? %+v\n", c)
	for i := 0; i < c; i++ {
		go fmt.Println("Hello World..", i)
	}

	var s string
	fmt.Scanln(&s)
}
