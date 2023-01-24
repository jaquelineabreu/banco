package contas


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

func (c *ContaCorrente) Transferir(valoDaTransferencia float64, contaDestino *ContaCorrente)bool{
	if valoDaTransferencia < c.saldo && valoDaTransferencia > 0 {
		c.saldo -= valoDaTransferencia
		contaDestino.Depositar(valoDaTransferencia)
		return true
	}else{
		return false
	}
}
