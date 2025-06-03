package main

import (
	"fmt"
	"time"
)

func main(){
	const numeroTareas = 5
	tareas := make(chan int, numeroTareas)
	results := make(chan int, numeroTareas)

	for i := 1; i <= 3; i++ {
		go worker(i, tareas, results)
	}

	for p := 1; p <= numeroTareas; p++ {
		tareas <- p
	}
	close(tareas)

	for r := 1; r <= numeroTareas; r++ {
		fmt.Println("Resultados:", <- results)
	}

}

func worker(id int, tareas <- chan int, results chan <- int){
	for tarea := range tareas{
		fmt.Println("Worker ID:", id, "tarea iniciada", tarea)
		time.Sleep(time.Second * 2)

		fmt.Println("Worker ID:", id, "procesando tarea:", tarea)
		results <- tarea * 2
		fmt.Println("Worker ID:", id, "fin de tarea:", tarea)
	}
}