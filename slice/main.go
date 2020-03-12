package main

import "fmt"

func main() {
	// len(): # de elementos en el slice
	// cap(): # de elementos del array donde apunta el slice, a partir del índice de donde se creó el slice
	/* 	food := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}
	   	fruits := food[1:3] // "🍓", "🍋"
	   	fruits = append(fruits, "🍍", "🍈", "🍐")

	   	fmt.Println("food", food)
	   	fmt.Println("fruits", fruits)
	   	fmt.Println("tamaño", len(fruits))
	   	fmt.Println("capacidad", cap(fruits)) */

	//fruits := []string{"🍓", "🍋"}
	/* fruits := make([]string, 0, 3)

	fruits = append(fruits, "🍍", "🍓", "🍋", "🍎")
	fmt.Println("fruits", fruits)
	fmt.Println("tamaño", len(fruits))
	fmt.Println("capacidad", cap(fruits)) */

	fruits := []string{}
	fmt.Println("fruits", fruits)
	fmt.Println("tamaño", len(fruits))
	fmt.Println("capacidad", cap(fruits))
	fmt.Println(fruits == nil)
}
