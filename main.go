package main

import (

	"fmt"
	"github.com/jaquelineabreu/banco/contas"
	
) 

func PagarBoleto(conta verificarConta, valorDoBoleto float64){
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main(){	

	contaDaMarilia := contas.ContaPoupanca{}
	contaDaMarilia.Depositar(100)
	PagarBoleto(&contaDaMarilia, 60)

	fmt.Println(contaDaMarilia.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, -400)

	fmt.Println(contaDaLuiza.ObterSaldo())

	
}