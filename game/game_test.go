package game

import (
	"testing"
)

func TestLetterInWord(t *testing.T) {
	word := []string{"A", "n", "a", "i", "s"}
	guess := "n"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s", word, guess)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"A", "n", "a", "i", "s"}
	guess := "z"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contain letter %s", word, guess)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Word does not contain a letter. Should return an error")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	compareState(t, "goodGuess", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	compareState(t, "badGuess", g.State)
}

func TestGameWin(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	compareState(t, "won", g.State)
}

func compareState(t *testing.T, expectedState, currentState string) bool {
	if expectedState != currentState {
		t.Errorf("state must be equals to %sn, got %s", expectedState, currentState)
		return false
	}
	return true
}
