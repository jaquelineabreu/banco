package contas

import "github.com/jaquelineabreu/banco/clientes"

type contaPoupanca struct{
	Titular 		clientes.Titular
	NumeroAgencia 	int64
	NumeroConta 	int64
	Operacao 		int64
	saldo 			float64

}

func (c *contaPoupanca) Sacar(valordoSaque float64) string{
	podeSacar := valordoSaque > 0 && valordoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valordoSaque
		return "Saque realizado com sucesso"
	}else{
		return "saldo insuficiente"
	}
}

func (c *contaPoupanca) Depositar(valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}else{
		return "Valor deposito é menor que zero", c.saldo
	}

}

func (c *contaPoupanca) ObterSaldo() float64{
	return c.saldo
}
