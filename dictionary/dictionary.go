package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

func Load(filename string) error {
	file, err := os.Open(filename)
	if nil != err {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PickWord() string {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}
