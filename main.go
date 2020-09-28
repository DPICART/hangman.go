package main

import (
	"fmt"
	"os"

	"hangman.go/dictionary"
	"hangman.go/game"
)

func main() {

	err := dictionary.Load("wordslist.txt")
	if nil != err {
		fmt.Printf("Could not load dictionary: %v \n", err)
		os.Exit(1)
	}

	word := dictionary.PickWord()
	g, err := game.New(len(word), word)

	if nil != err {
		fmt.Printf("Could not create game: %v \n", err)
		os.Exit(1)
	}

	game.DrawWelcome()

	guess := ""
	for {
		game.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}
		l, err := game.ReadGuess()
		if nil != err {
			fmt.Printf("Failed to read from console: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
