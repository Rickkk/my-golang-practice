package main

import "fmt"

func main(){
	x:=[]int{123, 256, 12321, 126, 9012, 234,   343, 234, 2565, 676,  877, 54, 445, 2334,  454}
	y:=x[0:5]

	var total1, total2 int
	for _, value := range x {
		if(value % 2 == 0){
			total1 += value*y[1]
		}else{
			total2 += value*y[2]
		}
	}

 fmt.Println(total1)
 fmt.Println(total2)
}