package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello User!")
	var x string = "y"
	for x == "y"{
		
		eye := []string{"👁", "👀", "🫣", "😳", "😵‍💫", "🙄", "😶‍🌫️", "🥴", "😎", "🤓", "🧐", "🤩", "😭", "😢"}
		mouth := []string{"👄", "🫦", "😗", "😙", "😚", "😘", "😙", "😶", "😮", "😯", "😲", "😛", "😜", "🤪", "😝", "🤑", "🤭", "😷"}
		extras := []string{"🕶️", "👓", "🥽", "🥸", "🎩", "🧢", "👒", "⛑️", "🎭", "🩹", "🦯", "🎧", "💄", "👑", "💍", "📿"}

		rand.Seed(time.Now().UnixNano())
		reye := eye[rand.Intn(len(eye))]

		rand.Seed(time.Now().UnixNano())
		rmouth := mouth[rand.Intn(len(mouth))]

		rand.Seed(time.Now().UnixNano())
		rextras := extras[rand.Intn(len(extras))]
		
			fmt.Println("Choosing eyes...")
			time.Sleep(2*time.Second)
			fmt.Printf("\nCaught : %s",reye)
			fmt.Println("\nChoosing mouth...")
			time.Sleep(2*time.Second)
			fmt.Printf("\nCaught : %s",rmouth)
			fmt.Println("\nChossing accessories...")
			time.Sleep(2*time.Second)
			fmt.Printf("\n Caught : %s",rextras)
			fmt.Println("\nCombining all the emojis...")
			time.Sleep(2*time.Second)
			fmt.Printf("\nYou Cooked : %s%s%s",rextras,reye,rmouth)

			fmt.Println("\nWanna make another (y/n)")
			fmt.Scanln(&x)
		}

}


