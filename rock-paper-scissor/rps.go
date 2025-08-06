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
	var x string="y"
	for x == "y" {
		z := []string{"rock", "paper", "scissor"}
		fmt.Println("Chose the game schematics: \n 1. Player VS Player \n 2. PLayer VS Computer")
		fmt.Printf("-> ")
		var schem int
		fmt.Scanln(&schem)
		switch schem {
		case 2:
			rand.Seed(time.Now().UnixNano())
			comp := z[rand.Intn(len(z))]
			fmt.Println("Enter your choice (rock/paper/scissor): ")
			var choice string
			fmt.Scanln(&choice)
			choice = strings.ToLower(choice)
			if comp == choice {
				fmt.Printf("User : %s \n Computer : %s \n Match Draw!", choice, comp)
			} else if comp == "rock" && choice == "paper" {
				fmt.Printf("User : %s \n Computer : %s \n Player Won!", choice, comp)
			} else if comp == "paper" && choice == "rock" {
				fmt.Printf("User : %s \n Computer : %s \n Computer Won!", choice, comp)
			} else if comp == "rock" && choice == "scissor" {
				fmt.Printf("User : %s \n Computer : %s \n Computer Won!", choice, comp)
			} else if comp == "scissor" && choice == "rock" {
				fmt.Printf("User : %s \n Computer : %s \n Player Won!", choice, comp)
			} else if comp == "paper" && choice == "scissor" {
				fmt.Printf("User : %s \n Computer : %s \n Player Won!", choice, comp)
			} else if comp == "scissor" && choice == "paper" {
				fmt.Printf("User : %s \n Computer : %s \n Computer Won!", choice, comp)
			}
		case 1:
			fmt.Println("Player1 (rock/paper/scissor): ")
			var p1 string
			fmt.Scanln(&p1)
			fmt.Println("Player2 (rock/paper/scissor): ")
			var p2 string
			fmt.Scanln(&p2)
			p1 = strings.ToLower(p1)
			p2 = strings.ToLower(p2)
			if p1 == p2 {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Match Draw!", p1, p2)
			} else if p1 == "rock" && p2 == "paper" {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Player2 Won!", p1, p2)
			} else if p1 == "paper" && p2 == "rock" {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Player1 Won!", p1, p2)
			} else if p1 == "rock" && p2 == "scissor" {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Player1 Won!", p1, p2)
			} else if p1 == "scissor" && p2 == "rock" {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Player2 Won!", p1, p2)
			} else if p1 == "paper" && p2 == "scissor" {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Player2 Won!", p1, p2)
			} else if p1 == "scissor" && p2 == "paper" {
				fmt.Printf("Player1 : %s \n Player2 : %s \n Player1 Won!", p1, p2)
			}
		}
		fmt.Println("\nPlay again? (y/n):")
		fmt.Scanln(&x)
		x = strings.ToLower(x)
	}
		fmt.Println("Hope you enjoyed! :)")
}
