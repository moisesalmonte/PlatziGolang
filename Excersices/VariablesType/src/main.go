package main

import (
	"fmt"
	"math"
)

func main(){
	//declarando variables
	entero := 5
	var entero_p int = 10
	var int_default int
	int_default = entero + entero_p

	fmt.Println("Resultado", int_default)
	//constantes
	const MI_CONSTANTE int = 0
	fmt.Println("Constante:", MI_CONSTANTE)
	
	//Binarios
	num_binario := 0xA
	fmt.Printf("Numero Binario: %b \n", num_binario)
	fmt.Printf("Tipo de variable: %T \n", num_binario)

	//Variables decimales
	PI := math.Pi
	var diametro float64 = 3.0

	fmt.Println("PIx3 =", PI*diametro)
	fmt.Printf("Decimal 2 num. %.2f \n", PI*diametro)
	//Resta de decimales
	var altura float32 = 4.67
	var base float32 = 5.59
	fmt.Println("Resta Base-altura =", base - altura)

	//Comparacion de boolean
	b_1 := true
	b_2 := false

	fmt.Println("boolean 1 =", b_1)
	fmt.Println("boolean 2 =", b_2)
	
	//Cadena de texto
	var texto string = "Hello world!"
	fmt.Println(texto)
	fmt.Println("🤓")
	var caracter rune = 'A'
	var caracter2 uint8 = 'a'
	var caracter3 byte = 'B'
	fmt.Println("rune \t uint8 \t byte")
	fmt.Printf("%c-%d \t %c-%d \t %c-%d \n", caracter, caracter, caracter2, caracter2, caracter3, caracter3)
	

}

