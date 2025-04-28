package main
import "fmt"

type MyArray struct
{
	arr []int
	curIdx int
}

func(m*MyArray)SetSize(size int){
	m.arr = make([]int, size);
	m.curIdx = 0;
}

func(m*MyArray)Insert(val int){
	m.arr[m.curIdx]= val
	m.curIdx ++

}

func(m*MyArray)Fetch(idx int) int {
   return m.arr[idx];
}

func(m*MyArray)Search(val int) int {
	for j:=0;j<len(m.arr);j++ {
		if m.arr [j]== val {
			return j
		}
	}
	return -1
}


func(m*MyArray)Print(){
	var val string
	for i:=0; i<m.curIdx; i++ {
		val += fmt.Sprintf("%d",m.arr[i])
		val += " "
	}
   fmt.Println(val)
}
func main(){
 var a MyArray;
 a.SetSize(5);
 a.Insert(2)
 a.Insert(4)
 a.Insert(6)
 a.Insert(1)
 a.Insert(2)


 a.Print()

 s:= a.Fetch(3)
 y:= a.Search(1)

 fmt.Println(s)
 fmt.Println(y)

}



