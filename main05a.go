package main

import "fmt"

func main05a() {
	resultado := somaRetornoNormal(10, 10)
	fmt.Println(resultado)
}

//funcao com tipo de retorno normal
func somaRetornoNormal(a int, b int) int {
	return a + b
}
