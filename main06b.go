package main

import "fmt"

//estudo de ponteiros: nao importa onde vc vai nudar o valor da variavel se o enderecamento da memoria eh o mesmo vai refletir em todos os locais com o mesmo enderecamento.

func main06b() {

	variavel := 10
	fmt.Println("valor da 'variavel' = ", variavel)
	fmt.Println("endereco da memoria 'variavel' = ", &variavel)

	//passando para a funcao o endereco de memoria.
	abc(&variavel)
	//imprimindo o valor da variavel alterada.
	fmt.Println("valor da 'variavel' ALETARADO na funcao = ", variavel)
}

//a funcao possui como parametro um ponteiro do tipo int que recebe o enderecamento de memoria.
func abc(a *int) {
	fmt.Println("----------")
	fmt.Println("a = ponteiro = endereco da memoria de 'variavel' = ", a)
	//vai trazer o valor real do ponteiro conforme endereco da memoria.
	fmt.Println("a = ponteiro = valor do ponteiro = valor da 'variavel' = ", *a)

	fmt.Println("----------")
	*a = 200
	fmt.Println("a = ponteiro = valor do ponteiro = ", *a)
	fmt.Println("a = ponteiro = endereco da memoria = ", a)
	fmt.Println("endereco da memoria 'a' = ", &a)

}
