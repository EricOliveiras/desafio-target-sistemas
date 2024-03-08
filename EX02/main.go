package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número para verificar se pertence à sequência de Fibonacci: ")
	fmt.Scanln(&numero)

	if isFibonacci(numero) {
		fmt.Printf("%d pertence à sequência de Fibonacci.\n", numero)
	} else {
		fmt.Printf("%d não pertence à sequência de Fibonacci.\n", numero)
	}
}

func isFibonacci(num int) bool {
	a, b := 0, 1
	for b < num {
		a, b = b, a+b
	}
	return b == num
}
