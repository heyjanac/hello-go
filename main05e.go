package main

import "fmt"

func main05e() {
	//funcoes anonimas e funcoes dentro de funcoes: tenho uma funcao que retorna uma funcao que retorna um inteiro.
	resultado := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}

	//a impressao tem que retornar o resultado e a funcao (o resultado da funcao)
	fmt.Println(resultado(54, 54, 54, 54)())
}
