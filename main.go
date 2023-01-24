package main

import (
	"fmt"

	"github.com/jaquelineabreu/banco/contas"
) 
func main(){	

	contaDaIvete := contas.ContaCorrente{Titular: "Ivete", Saldo: 500}
	contaDoJorge := contas.ContaCorrente{Titular: "Jorge", Saldo: 100}

	status := contaDoJorge.Transferir(-200, &contaDaIvete)

	fmt.Println(status)
	fmt.Println(contaDaIvete)
	fmt.Println(contaDoJorge)


	
}