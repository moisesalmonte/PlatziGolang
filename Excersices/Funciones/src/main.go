package main

import (
	"fmt"
)

/*
	Para correr este paquete desbes escribir go run src/*.go
*/


func main(){

	potencia(8,3)
	salida := deletreo("Hola")
	fmt.Println(salida)

	fmt.Println("------------------")
	fmt.Println(suma(2,3,4,5,6,7,7))
	

	//hola()
	//var car rune = '😊'
	//only_numbers(car)
	//var tiny int8 = 23
	//only_numbers(tiny)
	detetar_tipo_dato([4]int{})
	fmt.Printf("%T \n", [...]int{1,2,3,3,4})
	only_numbers(2.3)
	only_numbers(2)
	only_numbers(1000)
	many_args(4,5,6,7,"hola", "mundo", true, false, 3.5, 6.7, int8(5), '👀')
	//fmt.Println("Es una cadena:", isString("Hola ----"))
}

func potencia(x, y int){
	exp := x
	for i := 1; i < y; i++ {
		exp *= x
	}

	fmt.Printf("%d^%d es: %d \n", x, y, exp)
}

func deletreo(txt string) string {
	outStr := ""
	for _, itm := range txt{
		outStr = fmt.Sprintf("%s %c", outStr, itm)
	}

	return outStr[1:]
}

/* 
	Ejemplo de funcion con argumentos variables
	La funcion podra recibir n argumentos
*/
func suma(args ...any) int {
	fmt.Println(args)
	return 0
}