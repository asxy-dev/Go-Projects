// code written by asxy 
// discord : asxy.dev
// day 10

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var x string = "y"
	for x == "y" {
		z := []string{"rock", "paper", "scissor"}
		fmt.Println("Choose the game mode:\n 1. Player vs Player\n 2. Player vs Computer")
		fmt.Print("-> ")
		var schem int
		fmt.Scanln(&schem)

		switch schem {
		case 2:
			rand.Seed(time.Now().UnixNano())
			comp := z[rand.Intn(len(z))]
			fmt.Print("Enter your name (leave blank for Player): ")
			var name string
			fmt.Scanln(&name)
			if name == "" {
				name = "Player"
			}
			fmt.Print("Enter your choice (rock/paper/scissor): ")
			var choice string
			fmt.Scanln(&choice)
			choice = strings.ToLower(choice)
			if comp == choice {
				fmt.Printf("%s: %s | Computer: %s\nResult: Draw 🤝\n", name, choice, comp)
			} else if comp == "rock" && choice == "paper" || comp == "paper" && choice == "scissor" || comp == "scissor" && choice == "rock" {
				fmt.Printf("%s: %s | Computer: %s\nResult: %s Won 🎉\n", name, choice, comp, name)
			} else {
				fmt.Printf("%s: %s | Computer: %s\nResult: Computer Won 💻\n", name, choice, comp)
			}

		case 1:
			fmt.Print("Enter number of players (2 or 3): ")
			var num int
			fmt.Scanln(&num)
			if num < 2 || num > 3 {
				fmt.Println("Only 2 or 3 players are allowed.")
				continue
			}

			names := make([]string, num)
			choices := make([]string, num)

			for i := 0; i < num; i++ {
				fmt.Printf("Enter name for Player %d (leave blank for default): ", i+1)
				fmt.Scanln(&names[i])
				if names[i] == "" {
					names[i] = fmt.Sprintf("Player %d", i+1)
				}
				fmt.Printf("%s's choice (rock/paper/scissor): ", names[i])
				fmt.Scanln(&choices[i])
				choices[i] = strings.ToLower(choices[i])
			}

			winner := ""
			draw := false

			if num == 2 {
				a, b := choices[0], choices[1]
				if a == b {
					draw = true
				} else if a == "rock" && b == "scissor" || a == "paper" && b == "rock" || a == "scissor" && b == "paper" {
					winner = names[0]
				} else {
					winner = names[1]
				}
			} else {
				a, b, c := choices[0], choices[1], choices[2]
				if a == b && b == c {
					draw = true
				} else if (a == "rock" && b == "scissor" && c == "scissor") ||
					(a == "paper" && b == "rock" && c == "rock") ||
					(a == "scissor" && b == "paper" && c == "paper") {
					winner = names[0]
				} else if (b == "rock" && a == "scissor" && c == "scissor") ||
					(b == "paper" && a == "rock" && c == "rock") ||
					(b == "scissor" && a == "paper" && c == "paper") {
					winner = names[1]
				} else if (c == "rock" && a == "scissor" && b == "scissor") ||
					(c == "paper" && a == "rock" && b == "rock") ||
					(c == "scissor" && a == "paper" && b == "paper") {
					winner = names[2]
				} else {
					draw = true
				}
			}

			for i := 0; i < num; i++ {
				fmt.Printf("%s: %s\n", names[i], choices[i])
			}

			if draw {
				fmt.Println("Result: Draw 🤝")
			} else {
				fmt.Printf("Result: %s Won 🎉\n", winner)
			}
		}
		fmt.Print("Play again? (y/n): ")
		fmt.Scanln(&x)
		x = strings.ToLower(x)
	}
	fmt.Println("Thanks for playing! 👋")
}
