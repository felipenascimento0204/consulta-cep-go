package main

import (
	"consulta-cep/postmon"
	"fmt"
)

const versao = 1.0

func main() {

	var cep string

	fmt.Println("Sistema de consulta de Endereço versão", versao)
	fmt.Println("")
	fmt.Println("Informe o cep que deseja consultar:")
	fmt.Scanln(&cep)

	postmon.ConsultarCep(cep)
}
