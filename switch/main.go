package main

import "fmt"

func main() {
	emoji := "🍎"
	switch {
	case emoji == "🐈" || emoji == "🐕":
		fmt.Println("es un animal")
	case emoji == "🍌" || emoji == "🍎":
		fmt.Println("es una fruta")
	default:
		fmt.Println("no es un animal ni una fruta")
	}
}
