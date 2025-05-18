package main

import (
	"fmt"
	"strconv"
)

func main(){
	var cadena string
	fmt.Println("How many pizza slices rest:")
	num := 0
	num = -1
	for{
		fmt.Scanln(&cadena)
		n, err := strconv.Atoi(cadena)
		if err == nil {
			num = n
			break
		}
		fmt.Println("Error in the input:", cadena)
	}
	
	boxes := num / 8
	slc := num % 8
	if slc > 4 {
		fmt.Printf("Boxes: %d, more than middle (slices): %d \n", boxes, slc)
	} else if slc < 4 && slc > 0{
		fmt.Printf("Boxes: %d, less than middle (slices): %d \n", boxes, slc)
	} else if slc == 4{
		fmt.Printf("Boxes: %d, and middle box (slices): %d \n", boxes, slc)
	} else {
		fmt.Printf("Boxes: %d completes \n")
	}
}