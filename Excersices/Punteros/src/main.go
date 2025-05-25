package main

import "fmt"

func main(){
	var miVariable int = 10
	ref := &miVariable
	*ref = 2
	fmt.Println(miVariable)
	referencia(&miVariable)
	fmt.Println(miVariable)

	fmt.Println("-----------")
	var puntero *int
	/* *puntero = 99 */
	fmt.Println(&puntero)

	//Uso de puntero por la libreria fmt
	fmt.Println("Escribe una palabra, luego la repito dos veces")
	var entrada string = ""
	fmt.Scanln(&entrada)
	fmt.Println(entrada, entrada)
}

func referencia(b *int){
	*b = 100
}