package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Input Text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	censor := []string{"gay", "fuck", "bitch", "dick", "pussy", "ass", "motherfucker", "mf" , "randi" , "muji" , "nigga"}

	for _, word := range censor {
		if strings.Contains(text, word) {
			text = strings.ReplaceAll(text, word, "*****")
		}
	}

	fmt.Println("Censored Text:", text)
}
