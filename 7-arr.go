package main 

import "fmt"


// поиск максимального значения в массиве

func main(){
	x := [5]float64{ 98, 93, 77, 82, 83 }
	var max float64 = 0
	for _, value := range x{
		if value > max {
			max = value
		}
	}
	fmt.Println(max);
}