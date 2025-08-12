package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello Chef!")
	var x string = "y"
	for x == "y" {
		food := []string{"🍕", "🍔", "🌭", "🍟", "🥗", "🍝", "🍣", "🍤", "🥪", "🥩", "🍛", "🥘", "🫔", "🍜", "🫓", "🥞"}
		drink := []string{"🥤", "🧃", "🍵", "☕", "🍺", "🍷", "🍹", "🧋", "🥛", "🍶", "🍼", "🧉"}
		dessert := []string{"🍩", "🍪", "🍫", "🍰", "🧁", "🥧", "🍦", "🍨", "🍮", "🍯", "🍬", "🍭"}
		extra := []string{"🌶️", "🧂", "🍯", "🧈", "🥬", "🫛", "🫑", "🥒", "🍄", "🧄", "🧅"}

		rand.Seed(time.Now().UnixNano())
		f := food[rand.Intn(len(food))]
		d := drink[rand.Intn(len(drink))]
		ds := dessert[rand.Intn(len(dessert))]
		e := extra[rand.Intn(len(extra))]

		fmt.Println("Choosing main dish...")
		time.Sleep(1 * time.Second)
		fmt.Println("Caught:", f)
		fmt.Println("Choosing drink...")
		time.Sleep(1 * time.Second)
		fmt.Println("Caught:", d)
		fmt.Println("Choosing dessert...")
		time.Sleep(1 * time.Second)
		fmt.Println("Caught:", ds)
		fmt.Println("Adding extras...")
		time.Sleep(1 * time.Second)
		fmt.Println("Caught:", e)
		fmt.Println("Serving your full course...")
		time.Sleep(2 * time.Second)
		fmt.Printf("Here is your meal: %s %s %s %s\n", f, d, ds, e)

		fmt.Println("Wanna cook again? (y/n)")
		fmt.Scanln(&x)

		
	}
	fmt.Println("Enjoy eating!")
}
