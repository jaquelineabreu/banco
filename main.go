package main

import (

	"fmt"
	"github.com/jaquelineabreu/banco/contas"
	
) 
func main(){	

	contaDaMarilia := contas.ContaPoupanca{}
	contaDaMarilia.Depositar(100)
	contaDaMarilia.Sacar(45)

	fmt.Println(contaDaMarilia.ObterSaldo())
	
}