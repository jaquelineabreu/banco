package main

import "fmt"

type ContaCorrente struct{
	titular string 
	numeroAgencia int64
	numeroConta int64
	saldo float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) string{
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso"
	}else{
		return "Saldo insuficiente"
	}
}

func main(){
	

	contaDaIvete := ContaCorrente{}

	contaDaIvete.titular = "Ivete"
	contaDaIvete.saldo = 500

	fmt.Println(contaDaIvete.saldo)

	fmt.Println(contaDaIvete.Sacar(-700))

	fmt.Println(contaDaIvete.saldo)
}