package main

import "fmt"

//estudo de ponteiros

func main06a() {

	//atribuindo valor 10 para a variavel a
	//a variavel 'a' é apontado para um lugar na memoria do computador
	//e este lugar na memoria vai guardar o valor 10
	//memoria-endereco(10) <------ a <------ 10
	//qnd quiser mudar o valor da variavel 'a' para 20, o 'a' vai chamar o endereco e o endereco tem o valor que sera atualizado.
	//memoria-endereco(20) <------ a <------ 20

	a := 10
	//para saber o endereco da memoria onde a variavel esta apontando eh soh colocar o '&' para impressao: fmt.Println(&a) = 0xc0000a2058 [endereco]
	//memoria-0xc0000a2058(10) <------ a <------ 10
	fmt.Println("valor da variavel 'a' = ", a)
	fmt.Println("endereco da memoria 'a' = ", &a)
	//o endereco da memoria fica catalogado nos ponteiros
	//a variavel ponteiro recebe o endereco de memoria
	var ponteiro *int = &a
	fmt.Println("ponteiro = ", ponteiro)
	//vai trazer o valor real do ponteiro conforme endereco da memoria.
	fmt.Println("valor do ponteiro = ", *ponteiro)

	fmt.Println("----------")
	*ponteiro = 50
	fmt.Println("valor do ponteiro = ", *ponteiro)
	fmt.Println("ponteiro = ", ponteiro)
	fmt.Println("endereco da memoria 'a' = ", &a)
	fmt.Println("valor da variavel 'a' = ", a)

	fmt.Println("----------")
	b := &a
	fmt.Println("valor da variavel 'b' = endereco da memoria 'a' = ", b)
	fmt.Println("endereco da memoria 'b' = ", &b)
	fmt.Println("valor do ponteiro 'b' = valor do ponteiro 'a' = ", *b)
	fmt.Println("ponteiro = ", ponteiro)
	fmt.Println("valor do ponteiro = ", *ponteiro)

	fmt.Println("----------Alterar o valor do ponteiro")
	*b = 60
	fmt.Println("valor da variavel 'b' = endereco da memoria 'a' = ", b)
	fmt.Println("endereco da memoria 'b' = ", &b)
	fmt.Println("valor do ponteiro 'b' diferente  valor do ponteiro 'a' = ", *b)
	fmt.Println("ponteiro = ", ponteiro)
	fmt.Println("valor do ponteiro = ", *ponteiro)
	//valor da variavel de a = 60 pq esta apontando para o mesmo endereco de memoria de b, a e ponteiro.
	fmt.Println("valor da variavel 'a' = ", a)

}
