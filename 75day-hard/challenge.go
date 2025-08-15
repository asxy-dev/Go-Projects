package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	day := 1
	streak := 0
	choice := 0
	for choice < 1 || choice > 3 {
		fmt.Println("The 75 Hard Challenge Begins!")
		fmt.Println("Choose difficulty: 1. Easy  2. Medium  3. Hard")
		fmt.Scan(&choice)
	}
	for day <= 75 {
		var tasks []string
		switch choice {
		case 1:
			tasks = []string{"Drink 2L water", "Walk 5k steps", "Read 5 pages"}
		case 2:
			tasks = []string{"Drink 3L water", "Workout 45 mins", "Read 10 pages", "No junk food", "Sleep 8 hours"}
		case 3:
			tasks = []string{"Drink 4L water", "Workout 90 mins", "Read 20 pages", "Follow strict diet", "Meditate 30 mins", "Cold shower", "No social media"}
		}
		fmt.Printf("\nDay %d: Complete %d tasks\n", day, len(tasks))
		completed := 0
		doneTasks := make([]bool, len(tasks))
		for completed < len(tasks) {
			fmt.Println("\nToday's tasks:")
			for i, t := range tasks {
				status := "[ ]"
				if doneTasks[i] {
					status = "[✔]"
				}
				fmt.Printf("%d. %s %s\n", i+1, status, t)
			}
			fmt.Print("Enter task number to mark as done: ")
			var num int
			fmt.Scan(&num)
			if num > 0 && num <= len(tasks) && !doneTasks[num-1] {
				doneTasks[num-1] = true
				completed++
				progress := int(float64(completed) / float64(len(tasks)) * 100)
				total := 50
				for p := 0; p <= progress; p++ {
					done := int(float64(p) / 100 * float64(total))
					fmt.Printf("\r[%s%s] %d%%", strings.Repeat("█", done), strings.Repeat(" ", total-done), p)
					time.Sleep(20 * time.Millisecond)
				}
				fmt.Println()
			}
		}
		streak++
		fmt.Printf("Day %d complete! Streak: %d\n", day, streak)
		day++
		time.Sleep(1 * time.Second)
	}
	fmt.Println("75 Hard Challenge Complete!")
}
