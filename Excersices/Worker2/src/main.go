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
	NUM_POKEMON 	= 151 // Primera generacion
	NUM_WORKERS 	= 5 // 1-9
	ESC 			= 27
	BLUE 			= "\x1b[34m"
	RED 			= "\x1b[31m"
	GREEN 			= "\x1b[32m"
	YELLOW 			= "\x1b[33m"
	RESET_COLOR 	= "\x1b[0m"
)

var isFirstPrint = true
var lines [NUM_WORKERS + 2]string

func main(){
	//Canal para enviar la tecla presionada
	key := make(chan rune)
	tty, _ := tty.Open() //Funcion para obtener las teclas que preciona el usuario en la terminal
	workerControls := make([]Controls, 0, NUM_WORKERS)
	for i := 1; i <= NUM_WORKERS; i++ {
		workerControls = append(workerControls, Controls{make(chan int), make(chan int), make(chan int), false})
	} 	
	
	pokeList := getPokemonList()
	pokeList_div := divPokeList(pokeList, NUM_WORKERS)
	for i := 0; i < len(workerControls); i++{
		go worker(&workerControls[i], pokeList_div[i], i+1, key)
	}

	go func(){// Rutina para enviar la tecla presionada al canal key
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
	
	//Ocultar cursor
	fmt.Print("\033[?25l")
	lines[len(lines) - 2], lines[len(lines) - 1] = menuOne()

	go func(){ //Rutina para dibujar la pantalla cada n milliseconds
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

	//Liberar el cursor al final del programa
	fmt.Print("\033[?25h")
	tty.Close()
}

type Controls struct {
	cancel chan int
	pause chan int
	resume chan int
	isDone bool
}

func worker(control *Controls, listURL []string, id int, k chan rune){
	for idx, url := range listURL	{
		select{
		case <- control.cancel:
			lines[id - 1] = fmt.Sprintf("%sRoutine %d, canceled by user%s", RED, id, RESET_COLOR)
			control.isDone = true
			k <- rune(1)
			return
		case <- control.pause:
			lines[id - 1] = fmt.Sprintf("%sRoutine %d, paused by user%s", YELLOW, id, RESET_COLOR)
			<- control.resume
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
			time.Sleep(1000 * time.Millisecond)
		}
	}
	lines[id - 1] = fmt.Sprintf("%sRoutine %d, is completed.%s", GREEN, id, RESET_COLOR)
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
		fmt.Printf("\033[%dA", len(txt)) //Subir el cursor n lineas
	}
	for _, line := range txt{
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
		action = "Pausar"
		txt[length -2],	txt[length -1] = menuTwo("⏸️"+action, length)
	}else if k == 'r' || k == 'R' {
		action = "Reanudar"
		txt[length -2],	txt[length -1] = menuTwo("▶️"+action, length)
	}else if k == 'c' || k == 'C' {
		action = "Cancelar"
		txt[length -2],	txt[length -1] = menuTwo("⏹️"+action, length)
	}else{
		return false, action
	}
	return true, action
}

func routineChoose(rutina rune, action string, txt *[len(lines)]string, cw []Controls) bool{
	length := len(txt)
	const DIV_K = 48
	num := rutina % DIV_K
	if num >= 1 && int(num) <= length - 2 {
		switch(action){
		case "Pausar":
			cw[num - 1].pause <- 1
		case "Reanudar":
			cw[num - 1].resume <- 1
		case "Cancelar":
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