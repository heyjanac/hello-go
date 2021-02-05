package main

import "fmt"

func main05d() {
	resultado := somaTudo(3, 5, 10, 4, 6, 343)
	fmt.Println(resultado)
}

//funcao variadica
func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}
	return resultado
}
