package main

import "fmt"

// проверка на четность числа

func half(l int) (int, bool){
	if l % 2 == 0 {
		return 1,true
	}else {
		return 0,false
	}
}

func max(args...int) (int){
    max := 0
    for _, v := range args {
        if v > max {
			max=v
		}
    }
    return max	
}


func makeOddGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
		i += 1
		if i % 2 != 0 {
			return
		}else {
			i += 1
			return
		}
        
    }
}

func fibonacci(numb uint) uint{
	if numb == 0 {
		// fmt.Println(0)
		return 0
	}
	if numb == 1 {
		// fmt.Println(1)
		return 1
	}

	result:= fibonacci(numb-1) + fibonacci(numb-2)
	// fmt.Println(result)
	return result
}

func main(){

	// var numb int
    // fmt.Print("Введите число: ")
    // fmt.Scanf("%d", &numb)	
	// fmt.Println(half(numb))


	// fmt.Println(max(21,213,213,2341,676,454,12,56,76,213,85854))



	// замыкание - генератор нечетных чисел
	// nextEven := makeOddGenerator()
	// i:=0
    // for i <= 100 {
    //     fmt.Println(nextEven()) 
    //     i = i + 1
	// }	

	// вычисление числа Фибоначчи
	var numb uint
    fmt.Print("Введите n-число Фибоначчи: ")
    fmt.Scanf("%d", &numb)	
	fmt.Println(fibonacci(numb))
}