// written by asxy
// discord : asxy.dev

package main

import (
	"fmt"
	"strings"
)

func main() {
	var pas string
	var wua string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&pas)
	runes := []rune(pas)

	l := len(runes)
	if l <= 7 || l >= 24 {
		fmt.Println("Password Must Be Greater Than 7 Characters And Shorter Than 24 Characters!")
	} else {
		score := 0
		strong := "*#@!$^&)(';.,></|][}{]1234567890"
		medium := "1234567890"
		easy := "QWERTYUIOPASDFGHJKLZXCVBNMabcdefghijklmnopqrstuvwxyz"
		for i := 0; i < l; i++ {
			if strings.Contains(strong, string(runes[i])) {
				score += 3
			} else if strings.Contains(medium, string(runes[i])) {
				score += 2
			} else if strings.Contains(easy, string(runes[i])) {
				score += 1
			}
		}
		if score >= l*2 {
			wua = "Strong Password"
		} else if score >= l {
			wua = "Medium Password"
		} else if score > 0 {
			wua = "Easy Password"
		} else {
			wua = "Super Easy || Tips: Change it ASAP !"
		}
		fmt.Println(wua)
	}
}
