// code written by asxy
// discord : asxy.dev
// day 15

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name   string
	HP     int
	Level  int
	Score  int
	MaxHP  int
}

type Monster struct {
	Name string
	HP   int
	Attack int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var name string
	fmt.Print("Enter your hero name: ")
	fmt.Scanln(&name)
	player := Player{Name: name, HP: 100, MaxHP: 100, Level: 1, Score: 0}
	difficulty := chooseDifficulty()
	monsters := generateMonsters(difficulty)
	fmt.Println("Get ready for battle!")
	time.Sleep(time.Second)

	for i, m := range monsters {
		fmt.Printf("\nMonster %d: %s appears! HP: %d\n", i+1, m.Name, m.HP)
		for m.HP > 0 && player.HP > 0 {
			fmt.Printf("\nYour HP: %d | Monster HP: %d\n", player.HP, m.HP)
			fmt.Println("Choose attack:\n1. Slash\n2. Fireball\n3. Heal")
			var choice int
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				dmg := rand.Intn(15) + 5
				m.HP -= dmg
				fmt.Printf("You slash %s for %d damage!\n", m.Name, dmg)
			case 2:
				dmg := rand.Intn(25) + 10
				m.HP -= dmg
				fmt.Printf("You cast fireball on %s for %d damage!\n", m.Name, dmg)
			case 3:
				heal := rand.Intn(20) + 10
				player.HP += heal
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				fmt.Printf("You healed for %d HP!\n", heal)
			default:
				fmt.Println("Invalid choice, you lose your turn!")
			}
			if m.HP <= 0 {
				fmt.Printf("%s defeated!\n", m.Name)
				player.Score += 50 * difficulty
				player.Level += 1
				player.HP += 10
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				break
			}
			mdmg := rand.Intn(m.Attack) + 5
			player.HP -= mdmg
			fmt.Printf("%s attacks you for %d damage!\n", m.Name, mdmg)
			if player.HP <= 0 {
				fmt.Println("You were defeated! Game Over!")
				fmt.Printf("Final Score: %d | Level reached: %d\n", player.Score, player.Level)
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("You defeated all monsters! Victory!")
	fmt.Printf("Final Score: %d | Level reached: %d\n", player.Score, player.Level)
}

func chooseDifficulty() int {
	fmt.Println("Choose difficulty:\n1. Easy\n2. Medium\n3. Hard")
	var d int
	fmt.Scanln(&d)
	switch d {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	default:
		fmt.Println("Invalid, defaulting to Easy")
		return 1
	}
}

func generateMonsters(difficulty int) []Monster {
	names := []string{"Goblin", "Orc", "Troll", "Wolf", "Zombie", "Slime", "Vampire", "Bandit", "Dragonling", "Skeleton"}
	var monsters []Monster
	for i := 0; i < 10; i++ {
		hp := rand.Intn(50*difficulty) + 50
		attack := rand.Intn(15*difficulty) + 5
		monsters = append(monsters, Monster{Name: names[rand.Intn(len(names))], HP: hp, Attack: attack})
	}
	return monsters
}
