package main

import (
	"fmt"

	"github.com/jaquelineabreu/banco/contas"
)

func main(){	

	contaDaIvete := contas.ContaCorrente{titular:"Ivete", saldo: 500}
	contaDoJorge := contas.ContaCorrente{titular: "Jorge", saldo: 100}

	status := contaDoJorge.Transferir(-200, &contaDaIvete)

	fmt.Println(status)
	fmt.Println(contaDaIvete)
	fmt.Println(contaDoJorge)


	
}