package main


import ("fmt"
 "time")

func Sleep(sec int, c chan string){
	select {
	case <-time.After(time.Duration(sec)*time.Second):
		fmt.Println(time.Now().Format(time.RFC3339))
		fmt.Println("timed out")
		c <- "ping"
	}
}

func main(){
	fmt.Println("Введите количество секунд для таймера")
	var dur int
	fmt.Scanln(&dur)
	fmt.Println(time.Now().Format(time.RFC3339))
	var c chan string = make(chan string)
	dur = 2
	go Sleep(dur, c)
	for{
		select 
		{
			case <-c:
			fmt.Println("Сообщение получено")
			return			
		}
	}
    
}