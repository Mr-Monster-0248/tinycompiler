package lexer

import (
	"bufio"
	"fmt"
	"os"
)

type Lexer struct {
	currChar byte
	currPos  int
}

func Parse(file *os.File) []Token {
	scanner := bufio.NewScanner(file)
	tokens := []Token{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		for _, char := range line {

		}

	}

	return tokens
}
