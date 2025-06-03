package main

import (
	"fmt"
	"errors"
)

func main(){
	err := errors.New("Mi Error de prueba")

	fmt.Println(err.Error())
	fmt.Printf("%T \n", err)
	n, err := Divede(90, 6, 5)
	fmt.Println(err)
	fmt.Printf("%d \n", n)
}

var DIVISION_ZERO = errors.New("Invalid division by zero")
var EMPTY_ARGS = errors.New("Args empty")

func Divede (nums ...int) (int, error) {
	if nums != nil{
		res := nums[0]
		for i := 1; i < len(nums); i++ {
			if (res == 0 || nums[i] == 0){
				return 0, DIVISION_ZERO
			}
			res /= nums[i]
		}
		return res, nil
	}

	return 0, EMPTY_ARGS
}