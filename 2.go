package main

import "fmt"

var bookName ="Теоретический минимум по Computer Science"

func main(){
	fmt.Println("1+1=", 1+1);
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")	
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(32132*42452)
	var x string = "Kirill Korobkov  is author of this program"
	fmt.Println(x)
	y := "Max"
	fmt.Println("My dog's name is", y)	
	fmt.Println(bookName)

	var(
		priceX = 100.0
		priceY = 120.0
		priceZ = 200.0
	)
	result := priceX+priceY+priceZ
	fmt.Println("Result is ")
	fmt.Println((result - result*0.2))
}