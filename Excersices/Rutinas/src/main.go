package main

import "fmt"

func main() {
    canal := make(chan string)
    go func(){canal <- "Hola mundo"}()
    msg := <- canal
    fmt.Println(msg)
    go func(){canal <- "Hola mundo 2"}()
    msg = <- canal
    fmt.Println(msg)
    go func(){canal <- "Hola mundo 3"}()
    msg = <- canal
    fmt.Println(msg)
}