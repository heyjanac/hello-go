package main

import "fmt"

//estudo de ponteiros: nao importa onde vc vai nudar o valor da variavel se o enderecamento da memoria eh o mesmo vai refletir em todos os locais com o mesmo enderecamento.

type Carro struct {
	Name string
}

//dentro do escopo desta funcao o nome foi alterado... MAS se o parametro for um ponteiro(*) ira alterar para todos os escopos.
func (c *Carro) andou() {
	c.Name = "Sentra"
	fmt.Println(c.Name, "andou")
}

//a impressao do nome do carro serah referente a este escopo do main.
func main() {
	carro := Carro{
		Name: "Ka",
	}
	carro.andou()
	fmt.Println(carro.Name)
}
