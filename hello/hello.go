package main

import (
	"fmt"
	"os"
)

func main() {

	exibirIntroducao()
	exibirMenu()
	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibirIntroducao() {
	nome := "Aline"
	versao := 1.1

	fmt.Println("Olá Sra.", nome)
	fmt.Println("Este programa esta na versão:", versao)

	fmt.Println()
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O Comando escolhido foi", comandoLido)

	return comandoLido
}

func exibirMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}
