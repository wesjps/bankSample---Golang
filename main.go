package main

import (
	cliente "bank/clientes"
	contas "bank/contas"
	"fmt"
)

func main() {

	clienteWes := cliente.Titular{Nome: "Wesley", CPF: "11122233344", Profissao: "Engenheiro de Software"}
	clienteJose := cliente.Titular{Nome: "José", CPF: "00011122233", Profissao: "Arquiteto de Software"}

	ctaDoWes := contas.ContaCorrente{}
	ctaDoWes.Titular = clienteWes
	ctaDoWes.NumeroAgencia = 2995
	ctaDoWes.NumeroConta = 3801
	ctaDoWes.Depositar(1000)

	ctaDoJose := contas.ContaCorrente{}
	ctaDoJose.Titular = clienteJose
	ctaDoJose.NumeroAgencia = 2995
	ctaDoJose.NumeroConta = 3801
	ctaDoJose.Depositar(1000)

	fmt.Println(ctaDoWes)
	fmt.Println(ctaDoJose)
	ctaDoWes.Transferir(100, &ctaDoJose)
	fmt.Println(ctaDoWes.ObterSaldo())
	fmt.Println(ctaDoJose.ObterSaldo())

}
