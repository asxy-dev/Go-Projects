// code written by asxy\
// discord: asxy.dev

package main

import (
	"fmt"
	"strings"
)

func main() {
	var user string
	var clm string
	fmt.Println("Enter your username: ")
	fmt.Scanln(&user)
	runes := []rune(user)

	l := len(user)
	if l <= 4 {
		fmt.Println("Username must be greater than 4 characters.")
		return
	}

	valid := "*#@!$^&)(';.,></|][}[]])-"

	
	if strings.Contains(valid, string(runes[0])) {
		clm = "Username should start with a letter."
	} else {
		
		clm = "Username has been updated."
		for i := 0; i < l; i++ {
			if strings.Contains(valid, string(runes[i])) {
				clm = "Username can't contain any special symbols."
				break
			}
		}
	}
	fmt.Println(clm)
}