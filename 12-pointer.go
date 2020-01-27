package main

import "fmt"

func zero(x *int){
	*x = 12
}

func main(){
//  x:=10
//  fmt.Println(x)
//  zero(&x)
//  fmt.Println(x)

//  x := 1.5
//  square(&x)
//  fmt.Println(x)
a:=12
b:=1273
fmt.Println(a, b)
prac1(&a, &b)
fmt.Println(a, b)
}

func square(x *float64) {
    *x = *x * *x
}

func prac1(a *int, b*int){
	z:=*a
	*a=*b
	*b=z
	// b=z
	
}
