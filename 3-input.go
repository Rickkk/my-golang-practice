package main 

import "fmt"

func main(){
	fmt.Println("Введите  число:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 4
	fmt.Println(output)
}