package game

import (
	"fmt"
	"strings"
)

type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("Word %s must be at least 3 characters long", word)
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.AddError(guess)
		if g.TurnsLeft == 0 {
			g.State = "lost"
		}
	}
}

func (g *Game) AddError(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft = g.TurnsLeft - 1
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, c := range g.Letters {
		if c == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
