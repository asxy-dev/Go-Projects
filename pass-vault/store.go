//code written by asxy
//discord : asxy.dev
//day 12

package main
import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
)
func main(){
	fmt.Println("Hello User!")
	time.Sleep(2*time.Second)
	var savedPass []string
	var savedWeb []string
	var x string = "y"
	for x=="y"{
		fmt.Println("Please choose the schematics: \n 1.Add Password \n 2.View Password \n 3.Remove all \n 4. Quit")
		var schem int
		fmt.Scanln(&schem)
		switch (schem){
		case 1:
			fmt.Println("Input number of password you want to save: ")
			var times int
			fmt.Scanln(&times)
			for i:=0;i<times;i++{
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Website: ")
				web, _ := reader.ReadString('\n')
				web = strings.TrimSpace(web)
				savedWeb = append(savedWeb, web)
				fmt.Print("Password: ")
				pass, _ := reader.ReadString('\n')
				pass = strings.TrimSpace(pass)
				savedPass = append(savedPass, pass)
				fmt.Println("Saving...")
				time.Sleep(3*time.Second)
			}
		case 2:
			fmt.Println("Your saved informations are: ")
			for j , t := range savedWeb{
				fmt.Println(t, "-", savedPass[j])
			}
		case 3:
			fmt.Println("Removing...")
			time.Sleep(3*time.Second)
			savedPass = nil
			savedWeb = nil
			fmt.Println("Removed all the saved passwords")
		case 4:
			x="n"
		default:
			fmt.Println("Error caught while choosing schematics. Please try again later!")
		}
		if x=="n"{
			break
		}
		fmt.Print("Wanna add more (y/n): ")
		fmt.Scanln(&x)
	}
	var save string
	fmt.Println("\nDo you want to save the tasks? (y/n)")
	fmt.Scanln(&save)
	if strings.ToLower(save) == "y" && len(savedPass) > 0 {
		file, err := os.Create("passwords.txt")
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}
		defer file.Close()
		for i := range savedPass {
			file.WriteString("Website: " + savedWeb[i] + "\nPassword: " + savedPass[i] + "\n\n")
		}
		fmt.Println("Passwords saved successfully to passwords.txt")
	} else {
		fmt.Println("Passwords not saved.")
	}
}
