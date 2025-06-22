package main

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"os"
	"io" 
	"github.com/mattn/go-tty"
)

const (
	DOWNLOAD_PATH = "download/"
	NUM_POKEMON = 151 // Primera generacion
	NUM_WORKERS = 4 // 1-9
	ESC = 27
)
var isFirstPrint = true
var lines [NUM_WORKERS + 2]string

func main(){
	controls_workers := make([]Controls, 0, NUM_WORKERS)
	for i := 1; i <= NUM_WORKERS; i++ {
		controls_workers = append(controls_workers, Controls{make(chan int), make(chan int), make(chan int)})
	} 	
	
	lista := get_pokemon_list_gen_1()
	lista_div := dividir_lista(lista, 4)

	for index, wrk := range controls_workers{
		go worker(wrk, lista_div[index], index+1)
	}

	//Canal para enviar la tecla presionada
	key := make(chan rune)
	tty, _ := tty.Open() //Funcion para obtener las teclas que preciona el usuario en la terminal
    defer tty.Close()

	go func(){// Rutina para enviar la tecla presionada al canal key
		for {
			err := keyPress(key, tty)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
	
	//Ocultar cursor
	fmt.Print("\033[?25l")
	lines[len(lines) - 2], lines[len(lines) - 1] = menu_1()

	go func(){ //Rutina para dibujar la pantalla cada n milliseconds
		for{
			imprimirTextos(lines)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	isAccionPress := false
	accion := ""
	var kPress rune
	for {
		kPress = <- key
		if(isAccionPress){
			rutinaElegida(kPress, accion, lines, controls_workers)
		}else{
			if(kPress == ESC) { //Salir del programa
				return
			}else{
				isAccionPress, accion = accionElegida(kPress, lines)
			}
		}
	}

	//Liberar el cursor al final del programa
	fmt.Print("\033[?25h")
}

type Controls struct {
	cancel chan int
	pause chan int
	resume chan int
}

func worker(control Controls, list_url []string, id int){
	for idx, url := range list_url	{
		select{
		case <- control.cancel:
			lines[id - 1] = fmt.Sprintf("Rutina #%d, cancelada por el usuario", id)
			return
		case <- control.pause:
			lines[id - 1] = fmt.Sprintf("Rutina #%d, pausada por el usuario", id)
			<- control.resume
		default:
			res, err := descargar_archivo(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			name_file := strings.Split(url, "/")[8]
			err = guardar_archivo(res, name_file)
			if(err != nil){
				fmt.Println(err)
				return
			}
			lines[id - 1] = fmt.Sprintf("Rutina #%d, descargando %d de %d", id, idx+1, len(list_url))
			time.Sleep(1000 * time.Millisecond)
		}
	}
	lines[id - 1] = fmt.Sprintf("Rutina #%d, finalizada.", id)
}

func get_pokemon_list_gen_1() []string {
	var list_pokemon []string
	url := "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"
	
	for i := 1; i <= NUM_POKEMON; i++ {
		url_pokemon := fmt.Sprintf("%s%d.png", url, i)
		list_pokemon = append(list_pokemon, url_pokemon)
	}
	return list_pokemon
}


func descargar_archivo(url string) (*http.Response, error) {
	/* var response *http.Response
	err := errors.New("") */
	response, err := http.Get(url)
	if err != nil {
		return response, err
	}
	return response, nil
}

func guardar_archivo(res *http.Response, filename string) error {
	//creando carpeta
	err := os.MkdirAll(DOWNLOAD_PATH, 0777)
	if err != nil {
		return err
	}

	d_file, err2 := os.Create(DOWNLOAD_PATH + filename)
	if err != nil {
		return err2
	}

	_, err3 := io.Copy(d_file, res.Body)
	res.Body.Close()
	d_file.Close()

	return err3
} 

func dividir_lista(arreglo []string, dvd int) [][]string {
	length := len(arreglo)
	len_grp := length / dvd

	salida_arr := [][]string{}
	for i := len_grp; i <= len_grp * (dvd - 1); i += len_grp {
		salida_arr = append(salida_arr, arreglo[i-len_grp:i])
	}
	salida_arr = append(salida_arr, arreglo[(len_grp * (dvd - 1)):])
	return salida_arr
}

func keyPress(canal chan rune, t *tty.TTY) error {
	key, err := t.ReadRune()
	if err != nil {
		return err
	}

	canal <- key
	return nil
}

func imprimirTextos(txt [len(lines)]string){
	if(!isFirstPrint){
		fmt.Printf("\033[%dA", len(txt)) //Subir el cursor n lineas
	}
	for _, line := range txt{
		fmt.Printf("\033[0K%s\n", line)
	}
	isFirstPrint = false
}

func menu_1() (string, string) {
	str1 := "Presiona una tecla para seleccionar una accion:"
	str2 := "[P] pausar, [R] reanudar, [C] cancelar, [Esc] Salir"
	return str1, str2
}

func menu_2(accion string, length int) (string, string) {
	str1 := fmt.Sprintf("Haz eligido <<%s>>:", accion)
	str2 := fmt.Sprintf("Elige de 1 al %d, ha <%s>> o [Esc] para volver.", length, accion)
	return str1, str2
}

func accionElegida(k rune, txt [len(lines)]string) (bool, string){
	length := len(txt)
	var accion string
	if k == 'p' || k == 'P' {
		accion = "Pausar"
		txt[length -2],	txt[length -1] = menu_2(accion, length)
		accion = "Pausar"
	}else if k == 'r' || k == 'R' {
		accion = "Reanudar"
		txt[length -2],	txt[length -1] = menu_2(accion, length)
		accion = "Reanudar"
	}else if k == 'c' || k == 'C' {
		accion = "Cancelar"
		txt[length -2],	txt[length -1] = menu_2(accion, length)
		
	}else{
		return false, accion
	}
	return true, accion
}

func rutinaElegida(rutina rune, accion string, txt [len(lines)]string, cw []Controls){
	length := len(txt)
	const DIV_K = 48
	num := rutina % DIV_K
	if num >= 1 && int(num) <= len(txt) {
		switch(accion){
		case "Pausar":
			cw[num - 1].pause <- 1
		case "Reanudar":
			cw[num - 1].resume <- 1
		case "Cancelar":
			cw[num - 1].cancel <- 1
		}
		txt[length - 2], txt[length - 1] = menu_1()
	}
}