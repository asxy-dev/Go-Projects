// code written by asxy
//discord : asxy.dev

package main
import (
	"fmt"
	"math/rand"
	"time" 
)
func main(){
	
	
	var n int
	fmt.Printf("Please Input the number of dice you want to roll: ")
	fmt.Scanln(&n)
	total := 0
	for i:=1;i<=n;i++{
		fmt.Println("🧩 | Rolling..")
		time.Sleep(2*time.Second)
		roll := rand.Intn(6)+1
		if (roll==6){
			fmt.Println("🎲 | Lucky Roll, You Rolled: ",roll)
		}else if (roll==1) {
			fmt.Println("🎲 | Guess what! You got the lowest number ",roll)
		}else {
		fmt.Println("🎲 | You rolled: ",roll)
		
		total += roll
		}
	}
	if (total==2){
		fmt.Println("🐍 | Got a Snake Eye, Badluck: ",total)
	}else if (total==7){
		fmt.Println("🍀 | Got a Lucky number: ",total)
	} else if (total==12){
		fmt.Println("🔥 | Double six! You are on fire: ",total)
	}else{
	fmt.Println("💫 | Total: ",total)
	}

}