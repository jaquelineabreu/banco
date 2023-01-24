package main

import (
	"fmt"

	"github.com/jaquelineabreu/banco/contas"
	
) 
func main(){	

	//"github.com/jaquelineabreu/banco/clientes"

	//clienteMarisa := clientes.Titular{"Marisa", "123.111.123.12", "Desenvolvedora"}	

	//contaDaMarisa := contas.ContaCorrente{clienteMarisa, 123, 123456, 100}

	contaMarisa := contas.ContaCorrente{}

	contaMarisa.Depositar(100)

	fmt.Println(contaMarisa.ObterSaldo())
	
}