package contas

import cliente "bank/clientes"

type ContaPoupanca struct {
	Titular                              cliente.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiante"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso"
	} else {
		return "Valor de depósito inválido"
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	podeTransferir := c.saldo > valorDaTransferencia && valorDaTransferencia > 0

	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
