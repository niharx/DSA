package main
import "fmt"


func findMax(arr []int)int{
	var maxElement int = arr[0];

	for  j:= 0; j<len(arr); j++ {
		if maxElement < arr[j] {
			maxElement = arr[j]
		}
	}
	return maxElement
}

func findMin(arr []int)int{
	var minElement int = arr[0];

	for j:= 0; j<len(arr); j++ {
		if minElement > arr[j] {
			minElement = arr[j]
		}
	}
	return minElement
}

func main(){
   var arr = []int{5,3,14,23,2}
   fmt.Println(findMin(arr))
   fmt.Println(findMax(arr))
}