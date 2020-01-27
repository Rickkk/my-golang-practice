package main

import "fmt"

func main(){
	var x[5]  int;
	x[0] = 1984
	x[1] = 1984
	x[2] = 1956
	x[3] = 1956
	x[4] = 2016


    var total float64 = 0
    for i := 0; i < len(x); i++ {
        total += float64(2020 - x[i])
	}
	fmt.Println("Средний возраст")
    fmt.Println(total / float64(len(x)))
	
}