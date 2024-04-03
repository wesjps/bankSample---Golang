package main

import (
	"bank/contas"
	"fmt"
)

func main() {
	ctaDoWes := contas.ContaCorrente{}
	ctaDoWes.Titular = "Wesley"
	ctaDoWes.NumeroAgencia = 2995
	ctaDoWes.NumeroConta = 3801
	ctaDoWes.Saldo = 1000

	ctaDoJose := contas.ContaCorrente{}
	ctaDoJose.Titular = "Jose"
	ctaDoJose.NumeroAgencia = 2995
	ctaDoJose.NumeroConta = 3801
	ctaDoJose.Saldo = 1000

	fmt.Println(ctaDoWes)
	fmt.Println(ctaDoJose)
	ctaDoWes.Transferir(100, &ctaDoJose)
	fmt.Println(ctaDoWes)
	fmt.Println(ctaDoJose)

}
