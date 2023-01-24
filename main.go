package main

import (
	"fmt"

	"github.com/jaquelineabreu/banco/contas"
	"github.com/jaquelineabreu/banco/clientes"
) 
func main(){	

	clienteMarisa := clientes.Titular{"Marisa", "123.111.123.12", "Desenvolvedora"}

	contaDaMarisa := contas.ContaCorrente{clienteMarisa, 123, 123456, 100}

	fmt.Println(contaDaMarisa)
	
}