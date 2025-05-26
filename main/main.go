package main

import (
	//resolver problema com git
	"fmt"

	"./home/gustashow/Downloads/2vol.go/contas"
	//importar clientes
)

func main() {
	// contaDoArthur := contas.ContaCorrente{Titular: clientes.titular{
	// 	Nome:      "Bruno",
	// 	CPF:       "123.123.123.12",
	// 	Profissao: "Desenvolvedor"},
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   001,
	// 	Saldo:         140.90}

	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "dev"}
	//ciar uma var que porta o Titular, e adiciona ele
	//na criaçao da struct conta corrente..
	contaBruna := contas.ContaCorrente{clienteBruno, 123, 1234444, 100.40}

	fmt.Println(contaBruna)
}
