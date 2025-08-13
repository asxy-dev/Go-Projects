// code written by asxy
//discord : asxy.dev
//day 14

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Asxy's ATM Simulator!")
	time.Sleep(2 * time.Second)
	var pin int
	var ah string
	fmt.Println("Setup your first name only : ")
	fmt.Scanln(&ah)
	fmt.Println("Set your 4 digit pin : ")
	fmt.Scanln(&pin)
	if pin > 9999 || pin < 1000 {
		fmt.Println("Pin can't be more or less than 4 digits.")
		return
	} else {
		fmt.Println("Successfully Created Your Account.")
	}
	var x string = "y"
	var amt int = 1000
	for x == "y" {
		fmt.Println("Please choose the following: ")
		fmt.Println("1.Check Balance \n2.Withdraw Money \n3.Deposit Money \n4.Customer Support \n5.Quit")
		var cho int
		fmt.Scanln(&cho)
		if cho == 5 {
			fmt.Println("Thank you for using the ATM!")
			break
		}
		if cho == 4 {
			fmt.Println("Phone Number : +977-9812345678 / +977-9876543210 \nE-mail : asxyATM@gmail.com \nWebsite : asxyatm.com")
			continue
		}
		var cpin int
		attempts := 0
		for cpin != pin {
			fmt.Println("Please enter your transaction pin : ")
			fmt.Scanln(&cpin)
			if cpin != pin {
				attempts++
				if attempts == 1 {
					fmt.Println("Incorrect pin, wait 30 seconds before retrying...")
					time.Sleep(30 * time.Second)
				} else if attempts == 2 {
					fmt.Println("Incorrect pin again, wait 1 minute before retrying...")
					time.Sleep(1 * time.Minute)
				} else if attempts == 3 {
					fmt.Println("Incorrect pin again, wait 2 minutes before retrying...")
					time.Sleep(2 * time.Minute)
				} else if attempts == 4 {
					fmt.Println("Incorrect pin again, wait 1 hour before retrying...")
					time.Sleep(1 * time.Hour)
				} else {
					fmt.Println("Oi! You forgot your pin, press 4 to contact Customer Support.")
					break
				}
			}
		}
		if cpin != pin && attempts >= 5 {
			continue
		}
		switch cho {
		case 1:
			fmt.Println("Your balance is : Rs", amt)
		case 2:
			fmt.Println("Enter the amount to withdraw: ")
			var wit int
			fmt.Scanln(&wit)
			if wit > amt {
				fmt.Println("Insufficient funds in your account.")
			} else if wit <= 0 {
				fmt.Println("Invalid amount.")
			} else {
				fmt.Println("Processing your transaction...")
				time.Sleep(2 * time.Second)
				fmt.Println("Successfully withdrawed : Rs", wit)
				amt -= wit
				fmt.Println("Remaining : Rs", amt)
			}
		case 3:
			var depo int
			for {
				fmt.Println("Enter the amount to deposit:")
				fmt.Scanln(&depo)
				if depo > 10000 {
					fmt.Println("The maximum amount to deposit is 10000.")
				} else if depo <= 0 {
					fmt.Println("Invalid amount.")
				} else {
					fmt.Println("Successfully deposited.")
					amt += depo
					fmt.Println("Current Balance : Rs", amt)
					break
				}
			}
		default:
			fmt.Println("Please choose from the list above.")
		}
		fmt.Println("Re-open ATM? (y/n)")
		fmt.Scanln(&x)
		if x == "n" {
			fmt.Println("Thank you for banking with us! \n- Sincerely Asxy Bank Pvt. Ltd.")
		}
	}
}
