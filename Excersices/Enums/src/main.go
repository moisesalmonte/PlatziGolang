package main

import (
	"fmt"
	"encoding/hex"
)

type TextField struct{
	text string
	textColor Color
}

func main(){
	txtField := TextField{"Estoy aprendiendo Go en Platzi!", Color(Green)}
	/* fmt.Println(txtField.textColor.HexToRGB())
	fmt.Println(txtField.textColor.HexToString()) */

	txtField.Paint()
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

var color_ansi = map[Color] string{
	Blue: 	"\x1b[34m",
	Red: 	"\x1b[31m",
	Green:	"\x1b[32m",
	Yellow:	"\x1b[33m",
	Black:	"\x1b[30m",
	White:	"\x1b[37m",
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

func (c Color) ColorAnsi() string{
	return color_ansi[c]
}

func (t TextField) Paint(){
	color_reset_ansi := "\x1b[0m"
	txt_len := len(t.text)
	printLines(txt_len)
	fmt.Println("|" + t.textColor.ColorAnsi() + t.text + color_reset_ansi + "|")
	printLines(txt_len)
}

func printLines(length int){
	for i := 0; i < length + 2; i++ {
		fmt.Printf("%s", "-")
	}
	fmt.Printf("\n")
}
