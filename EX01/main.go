package main

import "fmt"

func main() {
	indice := 13
	soma := 0
	k := 0

	for k < indice {
		k += 1
		soma += k
	}

	fmt.Println(soma)
}
