package main 

import "fmt"




func main(){
	x := [6]string{"a","b","c","d","e","f"}
	y:=x[2:5]
	fmt.Println(y);

	// поиск минимального значения в массиве

	a := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}		
	min:= a[0]
	for _,value := range a{
		if value < min {
			min = value;
		}
	}
	fmt.Println(min);
}