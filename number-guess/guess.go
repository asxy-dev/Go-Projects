// code written by asxy
// discord : asxy.dev
// day 12

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello User")
	time.Sleep(2 * time.Second)
	var x string = "y"
	var score int
	var level int = 1
	var hs int
	for x == "y" {
		attempt := 1
		fmt.Println("Please choose the difficulty: ")
		var schem int
		fmt.Println("1. Easy \n2. Medium \n3. Hard")
		fmt.Printf("-> ")
		fmt.Scanln(&schem)
		easy := rand.Intn(100) + 1
		medium := rand.Intn(500) + 1
		hard := rand.Intn(1000) + 1
		switch schem {
		case 1:
			var guess int
			for guess != easy {
				fmt.Println("Guess a number (1-100): ")
				fmt.Scanln(&guess)
				if guess == easy {
					fmt.Println("Correct Guess in attempt:", attempt)
				} else if guess >= easy-5 && guess <= easy+5 {
					fmt.Println("Very close!")
					attempt++
				} else {
					fmt.Println("Incorrect guess!")
					attempt++
				}
			}
			score += 3
			level++
			if score > hs {
				hs = score
			}
		case 2:
			var guess int
			for guess != medium {
				fmt.Println("Guess a number (1-500): ")
				fmt.Scanln(&guess)
				if guess == medium {
					fmt.Println("Correct Guess in attempt:", attempt)
				} else if guess >= medium-5 && guess <= medium+5 {
					fmt.Println("Very close!")
					attempt++
				} else {
					fmt.Println("Incorrect guess!")
					attempt++
				}
			}
			score += 5
			level += 2
			if score > hs {
				hs = score
			}
		case 3:
			tries := 10
			var guess int
			for tries > 0 {
				fmt.Printf("Guess a number (1-1000): \nYou got %d tries left!\n -> ", tries)
				fmt.Scanln(&guess)
				tries--
				if guess == hard {
					fmt.Println("Correct Guess in attempt:", attempt)
					break
				} else if guess >= hard-5 && guess <= hard+5 {
					fmt.Println("Very close!")
					attempt++
				} else {
					fmt.Println("Incorrect guess!")
					attempt++
				}
			}
			if guess == hard {
				score += 10
				level += 3
				if score > hs {
					hs = score
				}
			} else {
				fmt.Println("You are out of tries!")
			}
		}
		fmt.Println("Play again? (y/n)")
		fmt.Scanln(&x)
	}
	fmt.Println("Final Outcome:")
	fmt.Printf("Score: %d | Level: %d | High Score: %d\n", score, level, hs)
	fmt.Println("Thanks for playing!")
}
