package main

import (
	"fmt"
	sl "slices"
	)

func main(){
	//Array
	var arr1 [4]int
	arr1[0] = 100
	fmt.Printf("Tipo: %T \n", arr1)
	fmt.Println(arr1[0])
	fmt.Println(len(arr1)) //size array
	fmt.Println(cap(arr1)) //Capacity array

	print("\n")
	//Slices
	var slice1 []int
	slice1 = append(slice1, 12)
	slice1 = append(slice1, 122)
	slice1 = append(slice1, 1222)
	slice1 = append(slice1, 12222)
	slice1 = append(slice1, 122222)
	slice1 = append(slice1, 1222222)
	slice1 = append(slice1, 12222222)
	slice1 = append(slice1, 122222222)
	slice1 = append(slice1, 1222222222)

	fmt.Printf("Tipo: %T \n", slice1)
	fmt.Println(slice1)
	fmt.Println("Tam. Slice1:", len(slice1)) //Size slice
	fmt.Println("Cap. Slice1:", cap(slice1)) //Capacity slice
	print("\n")

	slice2 := make([]int, 0, 10)
	slice2 = append(slice2, 4444444444)
	slice2 = append(slice2, 444444444)
	slice2 = append(slice2, 44444444)
	slice2 = append(slice2, 4444444)
	slice2 = append(slice2, 444444)
	slice2 = append(slice2, 44444)
	slice2 = append(slice2, 4444)
	slice2 = append(slice2, 444)
	slice2 = append(slice2, 44)
	slice2 = append(slice2, 4)
	slice2 = append(slice2, 0)

	fmt.Println(slice2)
	fmt.Println(slice2[:2]) //Imprimir los dos primero valores
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	//funcines de slices
	slice3 := sl.Concat(slice1, slice2)
	fmt.Println(slice3)
	fmt.Println(sl.Min(slice2))
	fmt.Println("Esta el 4 en el slice:", sl.Contains(slice2, 4))
	fmt.Println("Esta el 1 en el slice:", sl.Contains(slice2, 1))
	fmt.Println("Index de 4444 en el slice", sl.Index(slice2, 4444))
	fmt.Println("Index de 3223 en el slice", sl.Index(slice2, 3223))

	slice4 := make([]int, 0, 50)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	for i := 0; i < 50; i++ {
		slice4 = append(slice4, i)
	}

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	//delete
	slice4 = sl.Delete(slice4, 25, 50)
	fmt.Println(slice4)
	fmt.Println(len(slice4))

	//borrar los pares
	slice4_odd := sl.DeleteFunc(slice4, func(num int) bool {
		return num % 2 == 0
	})

	fmt.Println(slice4_odd)
	fmt.Println(slice4)
}
