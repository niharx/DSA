package main
import "fmt"

func main(){
	var num int;
	fmt.Print("Give me a number: ")
	fmt.Scanln(&num)

	if num==0 {
     fmt.Print("0")
	 return;
	}

    var rem int
	binary:=""
	for num>0 {
	rem = num % 2
    binary = fmt.Sprintf("%d%s",rem,binary)
	num = num / 2;
	}

	fmt.Print(binary)

}


