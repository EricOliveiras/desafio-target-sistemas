package main

import "fmt"

func main() {
	var palavra string
	fmt.Print("Digite uma palavra para ser revertida: ")
	fmt.Scanln(&palavra)

	palavraRevertida := reverse(palavra)
	fmt.Printf("A palavra %s, revertida fica: %s\n", palavra, palavraRevertida)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
