package main

import "fmt"

func detetar_tipo_dato(dato any){
	switch v := dato.(type) { // dato.(type) solo se puede usar en una sentencia switch
	case int:
		fmt.Println("es un int:", v)
	case string:
		fmt.Println("es un string", v)
	case bool:
		fmt.Println("es un bool", v)
	case int32: //Rune, caracter
		fmt.Println("es un caracter", v)
	case float64:
		fmt.Println("es un decimal:", v)
	case [5]int:
		fmt.Println("es un slice:", v)
	default:
		fmt.Println("Otro tipo:", v)
	}
}

//Funcion para detectar si es una cadena
func isString(dato any) bool {
	_, ok := dato.(string)
	return ok
}

/*
	Interface para definir los tipos de datos aceptados
*/
type Numero interface{
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

/*
	Esta funcion generica solo acepta los numeros registrados en la interface
*/
func only_numbers[n Numero](num n){
	fmt.Println("Es un numero:", num)
}

func many_args(arr ...any){
	fmt.Printf("Index\t\tValor\t\tType\n")
	fmt.Printf("-----\t\t-----\t\t----\n")

	for idx, dt := range arr{
		_, test := dt.(int32) //detectando si es un caracter
		if test {
			fmt.Printf("%d\t\t%c \t\t%T \n", idx, dt, dt)
			continue
		}
		fmt.Printf("%d\t\t%v \t\t%T \n", idx, dt, dt)
	}
}
