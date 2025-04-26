package main

import (
	"fmt"
	"os"
	"tinycompiler/lexer"
)

func main() {
	fmt.Println("Compiling...")

	file, err := os.Open("./test_input/fibo.BAS")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	lexer.Parse(file)
}
