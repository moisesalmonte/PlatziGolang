package main

import (
	"fmt"
	"time"
)

func main(){
	canal_resultado := make(chan string, 1)
	go consulta("", canal_resultado)

	select{
	case msg := <- canal_resultado:
		fmt.Println(msg)
	case <- time.After(300 * time.Millisecond):
		fmt.Println("Tiempo limite superado")
	}
}

func consulta(sql string, result chan string){
	time.Sleep(time.Millisecond * 500)
	result <- "rows inserted"
}