package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("Go relase test Buddy!!!")
	if len(args) > 0 {
		fmt.Println("Ur arg : ", args[0])
		return
	}
	fmt.Println("No arg. Try giving the arg")
}
