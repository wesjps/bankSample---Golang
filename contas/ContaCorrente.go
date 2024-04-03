package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiante"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso"
	} else {
		return "Valor de depósito inválido"
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	podeTransferir := c.Saldo > valorDaTransferencia && valorDaTransferencia > 0

	if podeTransferir {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
