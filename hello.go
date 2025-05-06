package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	versao := 1.1
	idade := 24

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}
