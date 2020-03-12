package main

import "fmt"

func main() {
	division(10, 2)
	division(40, 3)
	division(12, 0)
	division(20, 4)
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperandome del panic", r)
		}
	}()

	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("🤦‍♀️")
	}
}
