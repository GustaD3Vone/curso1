package main

import (
	//resolver problema com git
	"fmt"

	"github.com/GustaD3Vone/2vol.go/clientes"
	"github.com/GustaD3Vone/2vol.go/contas"
	//importar clientes
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.sacar(valorDoBoleto)
}

// interface pega a func de todas as pastas
// e a utiliza verificando as duas automaticamente
type verificarConta interface {
	sacar(valor float64) string
}

func main() {
	contaDoArthur := contas.ContaCorrente{Titular: clientes.titular{
		Nome:      "Bruno",
		CPF:       "123.123.123.12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123,
		NumeroConta:   001,
		Saldo:         140.90}

	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "dev"}
	//ciar uma var que porta o Titular, e adiciona ele
	//na criaçao da struct conta corrente..
	contaBruna := contas.ContaCorrente{clienteBruno, 123, 1234444, 100.40}

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.deposito(100)
	fmt.Println(contaExemplo.ObterSaldo())
	contaDoMark := contas.ContaPoupanca{}
	fmt.Println(contaDoMark)
	fmt.Println(contaDoArthur)
	fmt.Println(contaBruna)
	PagarBoleto(&contaDoMark, 60)
}
