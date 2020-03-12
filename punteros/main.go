package main

import "fmt"

func main() {
	fruit := "🍌"
	var p *string
	p = &fruit
	*p = "🍍"
	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", p, p, *p)
}
