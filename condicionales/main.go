package main

import "fmt"

func main() {
	/* 	emoji := "🐈"
	   	if emoji == "🌵" {
	   		fmt.Println("es un cactus")
	   	} else if emoji == "😋" {
	   		fmt.Println("es una carita")
	   	} else {
	   		fmt.Printf("el emoji es: %s\n", emoji)
	   	} */

	if emoji := "😋"; emoji == "🌵" {
		fmt.Println("es un cactus")
	} else if emoji == "😋" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %s\n", emoji)
	}

	fmt.Println(emoji)
}
