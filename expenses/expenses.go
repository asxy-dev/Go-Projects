//code written by asxy
// discord : asxy.dev
//day-9


package main

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
)

func main() {
	var list int
	var item []string
	var price []float64
	var x string = "y"

	for x == "y" {
		fmt.Print("Enter number of items: ")
		fmt.Scanln(&list)

		for i := 1; i <= list; i++ {
			fmt.Println("Adding...")
			time.Sleep(2 * time.Second)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Item : ")
			input, _ := reader.ReadString('\n')
			itemStr := strings.TrimSpace(input)
			item = append(item, itemStr)

			fmt.Print("Price : ")
			var p float64
			fmt.Scanln(&p)
			price = append(price, p)
		}

		fmt.Println("Options: \n 1. View Expenses \n 2. Calculate Total \n 3. Remove All \n 4. Quit")
		var opt int
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			for j, t := range item {
				fmt.Printf("Item %d: %s | Price: %.2f\n", j+1, t, price[j])
			}
		case 2:
			var total float64
			for _, p := range price {
				total += p
			}
			fmt.Println("Total: ", total)
		case 3:
			fmt.Println("Removing...")
			time.Sleep(2 * time.Second)
			item = nil
			price = nil
			fmt.Println("Successfully removed all items!")
		case 4:
			fmt.Println("Thanks for using our service!")
			return
		default:
			fmt.Println("Error occurred while selection, please re-try in few seconds!")
		}

		fmt.Print("Do you want to add more items? (y/n): ")
		fmt.Scanln(&x)
	}
}
