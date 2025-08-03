// code written by asxy
// discord : asxy.dev
// day : 6

package main
import "fmt"
func main(){
	var a,b int
	fmt.Println("Available Patterns: \n 1.Right-angle triangle(*) \n 2.Number Pyramids \n 3.Centered Triangle")
	fmt.Println("Input the Pattern Number You Want To Input: ")
	fmt.Scanln(&a)
	fmt.Println("Input the length of pattern")
	fmt.Scanln(&b)
	switch a{
	case 1:
		for i:=1;i<=b;i++{
			for j:=1;j<=i;j++{
				fmt.Print("*")
			}
			fmt.Println()
		}
	case 2:
		for i:=1;i<=b;i++{
			for j:=1;j<=i;j++{
				fmt.Print(j)
			}
			fmt.Println()
		}
	case 3:
		for i:=1;i<=b;i++{
			for j:=1;j<=b-i;j++{
				fmt.Print(" ")
			}
			for k:=1;k<=2*i-1;k++{
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}
