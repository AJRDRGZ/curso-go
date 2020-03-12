package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello Gophers"))
	if err != nil {
		fmt.Printf("Ocurrió un error al escribir el archivo: %v", err)
		return
	}
}
