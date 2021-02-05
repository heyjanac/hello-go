package main

import "fmt"

func main05b() {
	resultado, str := somaRetornoMultiplo(10, 20)
	fmt.Println(resultado, str)
}

//funcao com tipo de retorno multiplo
func somaRetornoMultiplo(a int, b int) (int, string) {
	return a + b, "somou!!"
}
