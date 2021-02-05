package main

import "fmt"

//metodo golang
type Carro05f struct {
	Nome string
}

//relacionar a funcao ao metodo Carro (c Carro) = chamado bind
func (c Carro05f) andar() {
	fmt.Println(c.Nome, "andou")
}

func main05f() {

	carro := Carro05f{
		Nome: "Sentra",
	}

	carro.andar()

	//funcoes anonimas e funcoes dentro de funcoes
	//tenho uma funcao que retorna uma funcao que retorna um inteiro.
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
