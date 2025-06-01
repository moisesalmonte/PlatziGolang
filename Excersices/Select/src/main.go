package main

import (
	"fmt"
	"time"
)

func main(){
	canal1 := make(chan string)
	canal2 := make(chan string)
	canal3 := make(chan string)

	go saludo("Juan", canal1, 4)
	go saludo("Nico", canal2, 1)
	go saludo("Kelvin", canal3, 2)

	for i := 1; i <= 3; i++ {
		select{
		case saludo1 := <- canal1:
			fmt.Println("#1", saludo1)
		case saludo2 := <- canal2:
			fmt.Println("#2", saludo2)
		case saludo3 := <- canal3:
			fmt.Println("#3", saludo3)
		}
	}

}

func saludo(name string, canal chan string, s time.Duration){
	saludo := fmt.Sprintf("Welcome %s", name)
	time.Sleep(s * time.Millisecond)
	canal <- saludo
}