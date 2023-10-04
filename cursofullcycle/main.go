package main

import (
	"cursofullcycle/entity"
	"cursofullcycle/math"
	"fmt"
	"log"
)

func main() {

	resul, err := math.Sum(2, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", resul)
	fmt.Printf("\n%v", math.SumAll(1, 1, 1, 1, 2))

	var u entity.User
	u.New(1, "aj", 32)
	fmt.Printf("\n%v - %v - %d", u.ID, u.Nome, u.Idade)
}
