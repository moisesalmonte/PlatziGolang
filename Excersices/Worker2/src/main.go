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
	DOWNLOAD_PATH 	= "download/"
	NUM_POKEMON 	= 151 //Pokemon primera generacion
	NUM_WORKERS 	= 5 // 1-9
	ESC 			= 27
	BLUE 			= "\x1b[34m"
	RED 			= "\x1b[31m"
	GREEN 			= "\x1b[32m"
	YELLOW 			= "\x1b[33m"
	RESET_COLOR 	= "\x1b[0m"
)

// Variable para detectar la primera impresion en pantalla
var isFirstPrint = true
// Arreglo para imprimir en pantalla de manera organizada
var lines [NUM_WORKERS + 2]string 

func main(){
	key 			:= make(chan rune) //Canal para enviar la tecla presionada por el usuario
	tty, _ 			:= tty.Open() //Funcion para obtener las teclas que preciona el usuario en la terminal
	workerControls 	:= make([]Controls, 0, NUM_WORKERS) //Slice para agrupar los canales para controlar las rutinas
	
	/*
		Ciclo for para inicializar las estructuras Controls 
		en el Slice, con channels anonimos. 
	*/
	for i := 1; i <= NUM_WORKERS; i++ {
		workerControls = append(workerControls, Controls{make(chan int), make(chan int), make(chan int), false, false})
	} 	
	
	pokeList 		:= getPokemonList() 
	pokeList_div 	:= divPokeList(pokeList, NUM_WORKERS)
	
	// Ciclo for para lanzar las rutinas
	for i := 0; i < len(workerControls); i++{
		go worker(&workerControls[i], pokeList_div[i], i+1, key)
	}

	// Rutina para escuchar la teclas presionada por el usuario
	go func(){
		for {
			err := keyPress(key, tty)
			if err != nil {
				fmt.Println(err)
				break
			}

			if isAllWorkerDone(workerControls) {
				break
			}
		}
	}()
	
	/* 
		Imprime en consola el codigo ANSI
		para ocultar el cursor
	*/
	fmt.Print("\033[?25l")
	
	/*
		Agregando las cadenas al arreglo, que se usuaran
		para imprimir el menu en pantalla, en dos ultima 
		posicion del arreglo
	*/
	lines[len(lines) - 2], lines[len(lines) - 1] = menuOne()

	//Rutina para dibujar la pantalla cada n milliseconds
	go func(){ 
		for{
			printLines(lines)
			time.Sleep(50 * time.Millisecond)
			if isAllWorkerDone(workerControls){
				break
			}
		}
	}()

	isActionPress := false
	action := ""
	var kPress rune
	/*
		Ciclo for para controlar la opciones
		tomadas por el usuario
	*/
	for {
		kPress = <- key
		if(isActionPress){
			isActionPress = routineChoose(kPress, action, &lines, workerControls)
		}else if kPress == 1 {
			if isAllWorkerDone(workerControls){
				break
			}
		}else{
			if(kPress == ESC) { //Salir del programa
				workerControls[0].isDone = true
				break
			}else{
				isActionPress, action = actionChoose(kPress, &lines)
			}
		}
	}
	
	cleanMenu()
	/* 
		Imprime en consola el codigo ANSI
		para mostrar el cursor
	*/
	fmt.Print("\033[?25h")
	//Cerrando el funcion que escucha las teclas presionadas
	tty.Close()
}

// Struct para agrupar los channels para controlar las rutinas
type Controls struct {
	cancel chan int
	pause chan int
	resume chan int
	isDone bool
	isPause bool
}

//Funcion Worker, para descargar y guardar el archivo
func worker(control *Controls, listURL []string, id int, k chan rune){
	for idx, url := range listURL	{
		select{
		case <- control.cancel:
			lines[id - 1] = fmt.Sprintf("%sRoutine %d, canceled by user%s", RED, id, RESET_COLOR)
			control.isDone = true
			k <- rune(1) //Enviando un dato al canal, para finalice si todas las rutinas fueron completadas 
			return
		case <- control.pause:
			lines[id - 1] = fmt.Sprintf("%sRoutine %d, paused by user%s", YELLOW, id, RESET_COLOR)
			control.isPause = true
			<- control.resume
			control.isPause = false
		default:
			lines[id - 1] = fmt.Sprintf("%sRoutine %d, downloading%s %d/%d", BLUE, id, RESET_COLOR, idx+1, len(listURL))
			res, err := downloadFile(url)
			if err != nil {
				fmt.Println(err)
				control.isDone = true
				k <- rune(1)
				return
			}
			nameFile := strings.Split(url, "/")[8]
			err = saveFile(res, nameFile)
			if(err != nil){
				fmt.Println(err)
				control.isDone = true
				k <- rune(1)
				return
			}
			//tiempo de espera para que las descargas sean mas lentas.
			time.Sleep(1000 * time.Millisecond)
		}
	}
	lines[id - 1] = fmt.Sprintf("%sRoutine %d, is completed.%s", GREEN, id, RESET_COLOR)
	// Tiempo de espera para que esperar que cambie el texto en pantalla
	time.Sleep(100 * time.Millisecond)
	control.isDone = true
	k <- rune(1)
}

func getPokemonList() []string {
	var pokemonList []string
	url := "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/"
	
	for i := 1; i <= NUM_POKEMON; i++ {
		pokemonURL := fmt.Sprintf("%s%d.png", url, i)
		pokemonList = append(pokemonList, pokemonURL)
	}
	return pokemonList
}

func downloadFile(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return response, err
	}
	return response, nil
}

func saveFile(res *http.Response, filename string) error {
	err := os.MkdirAll(DOWNLOAD_PATH, 0777)
	if err != nil {
		return err
	}

	fileDown, err2 := os.Create(DOWNLOAD_PATH + filename)
	if err != nil {
		return err2
	}

	_, err3 := io.Copy(fileDown, res.Body)
	res.Body.Close()
	fileDown.Close()

	return err3
} 

func divPokeList(pokeArray []string, dvd int) [][]string {
	length := len(pokeArray)
	divPokeArray := length / dvd

	divArrayReturn := [][]string{}
	for i := divPokeArray; i <= divPokeArray * (dvd - 1); i += divPokeArray {
		divArrayReturn = append(divArrayReturn, pokeArray[i-divPokeArray:i])
	}
	divArrayReturn = append(divArrayReturn, pokeArray[(divPokeArray * (dvd - 1)):])
	return divArrayReturn
}

func keyPress(channel chan rune, t *tty.TTY) error {
	key, err := t.ReadRune()
	if err != nil {
		return err
	}

	channel <- key
	return nil
}

func printLines(txt [len(lines)]string){
	if(!isFirstPrint){
		//Subir el cursor N lineas
		fmt.Printf("\033[%dA", len(txt)) 
	}
	for _, line := range txt{
		// \033[0k, este codigo ANSI es para borrar la informacion de la linea
		fmt.Printf("\033[0K%s\n", line)
	}
	isFirstPrint = false
}

func menuOne() (string, string) {
	str1 := "🟩 Press a key to select an action:"
	str2 := "🟩 [P] Pause, [R] Resume, [C] Cancel, [Esc] Exit"
	return str1, str2
}

func menuTwo(action string, length int) (string, string) {
	str1 := fmt.Sprintf("🟦 You chose, %s:", action)
	str2 := fmt.Sprintf("🟦 Choose from 1 to %d, to %s or [Esc] to go back.", length - 2, action)
	return str1, str2
}

func actionChoose(k rune, txt *[len(lines)]string) (bool, string){
	length := len(txt)
	var action string
	if k == 'p' || k == 'P' {
		action = "Pause"
		txt[length -2],	txt[length -1] = menuTwo("⏸️"+action, length)
	}else if k == 'r' || k == 'R' {
		action = "Resume"
		txt[length -2],	txt[length -1] = menuTwo("▶️"+action, length)
	}else if k == 'c' || k == 'C' {
		action = "Cancel"
		txt[length -2],	txt[length -1] = menuTwo("⏹️"+action, length)
	}else{
		return false, action
	}
	return true, action
}

func routineChoose(numRoutineChoose rune, action string, txt *[len(lines)]string, cw []Controls) bool{
	length := len(txt)
	const DIV_K = 48
	num := numRoutineChoose % DIV_K
	if num >= 1 && int(num) <= length - 2 {
		switch(action){
		case "Pause":
			cw[num - 1].pause <- 1
		case "Resume":
			cw[num - 1].resume <- 1
		case "Cancel":
			if(cw[num - 1].isPause){
				cw[num - 1].resume <- 1
			}
			cw[num - 1].cancel <- 1
		}
		txt[length - 2], txt[length - 1] = menuOne()
		return false
	}
	return true
}

func isAllWorkerDone(ctrls []Controls) bool{
	count_wrk_end := 0
	for _, ctrl := range ctrls{
		if ctrl.isDone{
			count_wrk_end++
		}
	}
	return count_wrk_end == len(ctrls)
}

func cleanMenu(){
	fmt.Print("\033[1A\033[0K\033[1A\033[0K")
}