package main

import "fmt"

func main(){

	potencia(8,3)
	salida := deletreo("Hola")
	fmt.Println(salida)

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