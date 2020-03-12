package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	hello("Alexys", "Lozada")

	emoji := "🐕"
	change(&emoji)
	fmt.Println(emoji)

	total := sum(2, 3)
	fmt.Println(total)

	texto := "ALejanDRo"
	minusc, mayusc := convert(texto)
	fmt.Println(minusc, mayusc)

	content, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Println(string(content))

	result, err := division(10, 2)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Println(result)

	nums := []int{1, 4, 70, 5, 67, 90, 2}
	resultFilter := filter(nums, func(num int) bool {
		if num >= 10 {
			return true
		}
		return false
	})
	fmt.Println(resultFilter)

	x := hello2("Alejandro")("Rodriguez")
	fmt.Println(x)

	fmt.Println(sumv(1, 2))

	func(text string) {
		fmt.Println("Hello", text)
	}("Gophers")
}

func sumv(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func hello2(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + " " + text
	}
}

func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}

func division(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("no puedes dividir por cero")
		return
	}

	result = dividendo / divisor
	return
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func change(value *string) {
	*value = "😋"
}

func hello(firstName string, lastName string) {
	fmt.Printf("Hello %s %s", firstName, lastName)
}
