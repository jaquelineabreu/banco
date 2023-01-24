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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}else{
		return "Valor deposito é menor que zero", c.saldo
	}

}

func main(){
	

	contaDaIvete := ContaCorrente{}

	contaDaIvete.titular = "Ivete"
	contaDaIvete.saldo = 500

	fmt.Println(contaDaIvete.saldo)
	
	status, valor := contaDaIvete.Depositar(1000)

	fmt.Println(status,valor)


	
}