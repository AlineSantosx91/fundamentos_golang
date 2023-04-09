package main

import (
	"fmt"
)

func main() {

	nome := "Aline"
	versao := 1.1

	fmt.Println("Olá Sra.", nome)
	fmt.Println("Este programa esta na versão:", versao)

	fmt.Println()

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int

	fmt.Scan(&comando)

	fmt.Println("O Comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não conheço este comando")
	}

}
