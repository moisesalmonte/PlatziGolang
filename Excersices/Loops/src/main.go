package main

import "fmt"

func main(){
	for i := 0; i < 10; i++{
		fmt.Println(i*2)
	}

	//for styled while loop
	p := 105
	for p > 100 {
		fmt.Println(p)
		p--
	}

	//loop infi...
	m := 10000000
	for{
		if m < 9999990 {
			break
		}
		fmt.Println("loop infi.")
		m--
	}

	//for recorrer un slice
	slice1 := []string {"a", "b", "c", "d", "e"}
	for idx, itm := range slice1{
		fmt.Println("Item:", itm, "Index:", idx)
	}
}

