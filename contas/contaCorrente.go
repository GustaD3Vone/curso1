package contas

import ".2vol.go/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) saque(valorSaque float64) string {
	podesacar := valorSaque <= c.Saldo && valorSaque > 0
	if podesacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insificiente"
	}
}
func (c *ContaCorrente) deposito(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "deposito realizado com suceso, Saldo atual:", c.Saldo
	} else {
		return "valor de deposito insuficiente", c.Saldo
	}
}
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.deposito(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
