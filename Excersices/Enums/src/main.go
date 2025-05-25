package main

import (
	"fmt"
	"encoding/hex"
)

type TextField struct{
	Width int
	Height int
	Text string
	Color Color
}

func main(){
	fmt.Println(Blue)

	txtField := TextField{100, 100, "Hello world!", Color(Blue)}
	fmt.Println(txtField.Color.HexToRGB())
	fmt.Println(txtField.Color.HexToString())

	red := "\033[31m"
    reset := "\033[0m"
    fmt.Println(red + "This text is red" + reset)
}

// Nuevo tipo para contener el indice
type Color int

// Inicializacion del indice utilizando iota
const (
	Blue = iota 
	Red
	Green
	Yellow
	Black
	White
)

var color_hex = map[Color] string{
	Blue: 	"0000FF",
	Red: 	"FF0000",
	Green: 	"00FF00",
	Yellow: "FFFF00",
	Black: 	"000000",
	White: 	"FFFFFF",
}

func (c Color) HexToRGB() [3][]byte {
	var cadena = [3]string{color_hex[c][0:2], color_hex[c][2:4], color_hex[c][4:6]}
	var rgb [3][]byte
	for idx, str := range cadena{
		bt, err := hex.DecodeString(str)
		if err != nil{
			fmt.Println(err)
			break
		}
		rgb[idx] = bt
	}
	return rgb
}

func (c Color) HexToString() string {
	return color_hex[c]
}
