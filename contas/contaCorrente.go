package contas


type ContaCorrente struct{
	Titular string 
	NumeroAgencia int64
	NumeroConta int64
	Saldo float64
}

func (c *ContaCorrente) Sacar(valordoSaque float64) string{
	podeSacar := valordoSaque > 0 && valordoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valordoSaque
		return "Saque realizado com sucesso"
	}else{
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	}else{
		return "Valor deposito é menor que zero", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valoDaTransferencia float64, contaDestino *ContaCorrente)bool{
	if valoDaTransferencia < c.Saldo && valoDaTransferencia > 0 {
		c.Saldo -= valoDaTransferencia
		contaDestino.Depositar(valoDaTransferencia)
		return true
	}else{
		return false
	}
}
