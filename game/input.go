package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Enter a letter: ")
		guess, err = reader.ReadString('\n')
		if nil != err {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Invalid letter input (one character pls)\n")
			continue
		}
		valid = true
	}
	return guess, nil
}
