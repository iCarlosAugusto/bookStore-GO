package main

import "fmt"

type ContaCorrete struct {
	titular string
	numeroAgencia int
	numerConta int
	saldo int
}

func (c *ContaCorrete) sacar(valor int) {
	if(c.saldo < valor) {
		fmt.Println("Operação não autorizada, saldo insuficiente");
		return;
	}
	newValue := c.saldo - valor;
	c.saldo = newValue;
}

func main() {
	conta1 := ContaCorrete{
		titular: "Carlos",
		numeroAgencia: 001,
		numerConta: 0031,
		saldo: 3000,
	}

	conta1.sacar(4000);
	fmt.Println(conta1.saldo);

}
