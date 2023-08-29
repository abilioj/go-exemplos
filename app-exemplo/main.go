package main

import (
	"fmt"
)

type Car struct {
	Name  string
	Madel string
}

func main() {
	nome := "abilio.jose"
	fmt.Println(nome)

	result, err := soma(5, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	carro := Car{"fusca", "vw"}
	fmt.Println(carro.Name)
}

func soma(a, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("soma maior que 10")
	}
	return a + b, nil
}
