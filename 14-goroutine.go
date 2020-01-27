package main

import "fmt"

func f(){
	for i:=0; i<=10; i++{
		fmt.Println(i)
	}
}

func main(){
go f()
var input string
fmt.Scanln(&input)
}