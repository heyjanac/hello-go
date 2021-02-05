package main

import (
	"fmt"
	"log"
	"net/http"
)

func main04a() {
	//retorna um erro caso contenha erro na chamada
	res, err := http.Get("http://google.com.br")
	if err != nil {
		log.Fatal("Erro ao fazer comunicacao: " + err.Error())
	}
	fmt.Println(res.Header)
}
