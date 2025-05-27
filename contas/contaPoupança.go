package contas

import "github.com/GustaD3Vone/2vol.go/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) saque(valorSaque float64) string {
	podesacar := valorSaque <= c.saldo && valorSaque > 0
	if podesacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}
func (c *ContaPoupanca) deposito(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "deposito realizado com suceso, saldo atual:", c.saldo
	} else {
		return "valor de deposito insuficiente", c.saldo
	}
}
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
