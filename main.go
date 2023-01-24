package main

import (
	"fmt"

	"github.com/jaquelineabreu/banco/contas"
	"github.com/jaquelineabreu/banco/clientes"
) 
func main(){	

	contaDaMarisa := contas.ContaCorrente{Titular:clientes.Titular{
		Nome: "Marisa",
		CPF:"123.111.123.12",
		Profissao: "Desenvolvedora"},
		NumeroAgencia:123, NumeroConta:123456, Saldo:100}

	fmt.Println(contaDaMarisa)
	
}