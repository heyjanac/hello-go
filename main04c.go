package main

import (
	"errors"
	"fmt"
	"log"
)

func main04c() {

	res, err := somaComErro(10, 2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func somaComErro(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}
	//senao deu problema a aplicacao
	return res, nil
}
