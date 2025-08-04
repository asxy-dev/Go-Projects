// code written by asxy.dev
// discord: asxy.dev
// day 8

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Hello User!")
	time.Sleep(1 * time.Second)

	var datatype int
	var op string
	var x, y interface{}

	fmt.Println("Please Input the datatypes \n 1. int-int \n 2. float-float \n 3. float-int \n 4. int-float")
	fmt.Scanln(&datatype)

	switch datatype {
	case 1:
		var a, b int
		fmt.Println("Input first integer:")
		fmt.Scanln(&a)
		fmt.Println("Input second integer:")
		fmt.Scanln(&b)
		x, y = a, b
	case 2:
		var a, b float64
		fmt.Println("Input first float:")
		fmt.Scanln(&a)
		fmt.Println("Input second float:")
		fmt.Scanln(&b)
		x, y = a, b
	case 3:
		var a float64
		var b int
		fmt.Println("Input first float:")
		fmt.Scanln(&a)
		fmt.Println("Input second integer:")
		fmt.Scanln(&b)
		x, y = a, float64(b)
	case 4:
		var a int
		var b float64
		fmt.Println("Input first integer:")
		fmt.Scanln(&a)
		fmt.Println("Input second float:")
		fmt.Scanln(&b)
		x, y = float64(a), b
	default:
		fmt.Println("Invalid selection")
		return
	}

	fmt.Println("Input the operator: (+,-,*,%,/,^)")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Println(x.(float64) + y.(float64))
	case "-":
		fmt.Println(x.(float64) - y.(float64))
	case "*":
		fmt.Println(x.(float64) * y.(float64))
	case "/":
		if y.(float64) != 0 {
			fmt.Println(x.(float64) / y.(float64))
		} else {
			fmt.Println("Division by zero")
		}
	case "%":
		if datatype == 1 {
			fmt.Println(int(x.(int)) % int(y.(int)))
		} else {
			fmt.Println("Modulo only works for int-int")
		}
	case "^":
		fmt.Println(math.Pow(x.(float64), y.(float64)))
	default:
		fmt.Println("Invalid operator")
	}
}
