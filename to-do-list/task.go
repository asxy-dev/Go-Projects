//code written asxy
//discord : asxy.dev
// day : 7

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var amt int
	var tasks []string

	fmt.Println("Hello User!")
	time.Sleep(1 * time.Second)
	fmt.Println("Please input the number of tasks: ")
	fmt.Scanln(&amt)

	reader := bufio.NewReader(os.Stdin)
	for i := 1; i <= amt; i++ {
		fmt.Printf("Input task No.%d : ", i)
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		tasks = append(tasks, task)
		fmt.Println("Saving...")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n1. View Task \n2. Remove Tasks \n3. Quit")
	var opt int
	fmt.Scanln(&opt)

	switch opt {
	case 1:
		fmt.Println("\nYour Tasks:")
		for j, t := range tasks {
			fmt.Printf("Task %d: %s\n", j+1, t)
		}
	case 2:
		tasks = []string{}
		fmt.Println("All tasks have been cleared!")
	case 3:
		fmt.Println("Thank you!")
	}

	var save string
	fmt.Println("\nDo you want to save the tasks? (y/n)")
	fmt.Scanln(&save)

	if strings.ToLower(save) == "y" && len(tasks) > 0 {
		file, err := os.Create("tasks.txt")
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}
		defer file.Close()

		for _, task := range tasks {
			file.WriteString(task + "\n")
		}
		fmt.Println("Tasks saved successfully to tasks.txt")
	} else {
		fmt.Println("Tasks not saved.")
	}
}
