package main

import "fmt"

func main05c() {
	resultado := somaRetornoNomeado(3, 5)
	fmt.Println(resultado)
}

////funcao com retorno nomeado
func somaRetornoNomeado(a int, b int) (result int) {
	result = a + b
	return
}
