package main
import "fmt"

func checkEvenOrOdd(){
	var num int;
	fmt.Println("give me a number for check")
	fmt.Scanln(&num)

if num&1 == 0 {
	fmt.Println("Number is even")
} else {
	fmt.Println("Number is odd")
}

}