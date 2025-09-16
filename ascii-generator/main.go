package main

import (
    "fmt"
    "github.com/common-nighthawk/go-figure"
)

func main() {
   var x string = "y"
   for x == "y" {
	fmt.Print("Enter text: ")
    var input string
    fmt.Scanln(&input)

    myFigure := figure.NewFigure(input, "", true)
    myFigure.Print()
	fmt.Println("Do you want to continue? (y/n)")
   	fmt.Scanln(&x)
	if (x!="y" && x!="n") {
		fmt.Println("Invalid input. Exiting.")
		break
	}
   }
   
}