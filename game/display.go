package game

import (
	"fmt"
)

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func DrawWelcome() {
	fmt.Println(`
		##################################
		############# HANGMAN ############
		##################################
	`)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
			    _____
			   |     |
			   |     o  
			   |    /|\
			   |    / \
			 _/|\_
			|     |________
			| RIP          |
			|______________|
		`
	case 1:
		draw = `
			    _____
			   |     |
			   |       
			   |    
			   |    
			 _/|\_
			|     |________
			|              |
			|______________|
		`
	case 2:
		draw = `
			    _____
			   |     
			   |       
			   |    
			   |    
			 _/|\_
			|     |________
			|              |
			|______________|
		`
	case 3:
		draw = `
				
			   |     
			   |       
			   |    
			   |    
			 _/|\_
			|     |________
			|              |
			|______________|
		`
	case 4:
		draw = `
				
			        
			          
			  
			       
			 _/ \_
			|     |________
			|              |
			|______________|
		`
	case 5:
		draw = `
				
			        
			          
			  
			       
			 
			       ________
			|              |
			|______________|
		`
	case 6:
		draw = `
				
			        
			          
			  
			       
			 
			       
			              
			______________
		`
	case 7:
		draw = `
				
			        
			          
			  
			       
			 
			       
			              
			_________
		`
	case 8:
		draw = `
				
			        
			          
			  
			       
			 
			       
			              
			   _
		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess\n")
	case "alreadyGuessed":
		fmt.Printf("Letter %s was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess %s is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost. The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You WON! The word was: ")
		drawLetters(g.Letters)

	}

}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
