package main

import "fmt"

//estudo de ponteiros: nao importa onde vc vai nudar o valor da variavel se o enderecamento da memoria eh o mesmo vai refletir em todos os locais com o mesmo enderecamento.

type Carro06c struct {
	Name string
}

//dentro do escopo desta funcao o nome foi alterado.
func (c Carro06c) andou06c() {
	c.Name = "Sentra"
	fmt.Println(c.Name, "andou")
}

//a impressao do nome do carro serah referente a este escopo do main.
func main06c() {
	carro := Carro06c{
		Name: "Ka",
	}
	carro.andou06c()
	fmt.Println(carro.Name)
}
