package main

import "fmt"

type student struct {
	Name  string
	Score []uint8
}

func main() {
	hello := "hello"
	// sports := map[string]string{"basketball": "🏀", "soccer": "⚽"}
	//nums := []uint8{2, 4, 6, 8}

	for _, v := range hello {
		fmt.Println(string(v))
	}
}
