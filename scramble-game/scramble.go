package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
)

func main() {
	var x string = "y"
	var score int = 0
	var level int = 1
	var hs int

	for x == "y" {
		fmt.Println("Please chose your game difficulty: \n 1. Easy \n 2. Medium \n 3. Hard")
		var schem int
		fmt.Scanln(&schem)
		easy := []string{"apple", "banana", "mango", "happy", "school","book","tree","cake","fish","ball","jump","girl","bird","lion","blue","hand","milk","star","play","door","rain","shoe","good","game","cold"}
		medium := []string{"giraffe", "picture", "holiday", "pencil", "window","animal","keyboard","bottle","circle","father","garden","house","jacket","kitten","letter","market","nature","orange","people","rabbit","school","sister","window","yellow","zebra"}
		hard := []string{"xylophone", "conspicuous", "suspisious", "academic", "pneumonia","absolute","butterfly","dangerous","elephant","hospital","identity","internet","language","mountain","newspaper","operator","question","sickness","triangle","umbrella","vacation","valuable","whispers","wildfire","adventure"}

		switch schem {
		case 1:
			var count int = 1
			for count <= 5 {
				rand.Seed(time.Now().UnixNano())
				choice := easy[rand.Intn(len(easy))]
				letters := strings.Split(choice, "")
				rand.Shuffle(len(letters), func(i, j int) {
					letters[i], letters[j] = letters[j], letters[i]
				})
				scrambled := strings.Join(letters, "")
				fmt.Println("Unscramble this word:", scrambled)
				var guess string
				fmt.Print("Your guess: ")
				fmt.Scanln(&guess)
				count++
				if strings.ToLower(guess) == choice {
					fmt.Println("Correct!")
					score += 3
					hs+=score
				} else {
					fmt.Println("Wrong! The word was:", choice)
				}
				level = (score / 50) + 1
				
				fmt.Println("Score:", score, "| Level:", level)
			}
		case 2:
			var count int = 1
			for count <= 5 {
				rand.Seed(time.Now().UnixNano())
				choice := medium[rand.Intn(len(medium))]
				letters := strings.Split(choice, "")
				rand.Shuffle(len(letters), func(i, j int) {
					letters[i], letters[j] = letters[j], letters[i]
				})
				scrambled := strings.Join(letters, "")
				fmt.Println("Unscramble this word:", scrambled)
				var guess string
				fmt.Print("Your guess: ")
				fmt.Scanln(&guess)
				count++
				if strings.ToLower(guess) == choice {
					fmt.Println("Correct!")
					score += 5
					hs+=score
				} else {
					fmt.Println("Wrong! The word was:", choice)
				}
				level = (score / 50) + 1
				
				fmt.Println("Score:", score, "| Level:", level)
			}
		case 3:
			var count int = 1
			for count <= 5 {
				rand.Seed(time.Now().UnixNano())
				choice := hard[rand.Intn(len(hard))]
				letters := strings.Split(choice, "")
				rand.Shuffle(len(letters), func(i, j int) {
					letters[i], letters[j] = letters[j], letters[i]
				})
				scrambled := strings.Join(letters, "")
				fmt.Println("Unscramble this word:", scrambled)
				var guess string
				fmt.Print("Your guess: ")
				fmt.Scanln(&guess)
				count++
				if strings.ToLower(guess) == choice {
					fmt.Println("Correct!")
					score += 10
					hs+=score
				} else {
					fmt.Println("Wrong! The word was:", choice)
				}
				
				level = (score / 50) + 1
				fmt.Println("Score:", score, "| Level:", level)
			}
		}
		fmt.Println("Play again? (y/n): ")
		fmt.Scanln(&x)
	}
	fmt.Println("Final Score:", score, "| Final Level:", level, "| High Score:",hs)
}
