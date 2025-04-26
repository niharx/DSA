package main
import "fmt"

func main(){
	num:=65535

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


