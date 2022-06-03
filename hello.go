package main

import (
	"fmt"
	"os"
)

/*
	PROJETO DE MONITORAMENTO DE SISTEMAS EM GOLANG
	Exemplo da aula introdutória de go
	Consegue:
		-Iniciar monitoramento dos sites;
		-Exibir logs dos sites;
*/
func main() {
	name := "Thiago" // inferência na string
	var version float32 = 0.3 // aqui é bom deixar, pq tem 2 tipos de float
	welcome(name,version)
}

// MOSTRA NOME E CHAMA O MENU
func welcome(name string, version float32) {
	fmt.Println(
		"Hello, Mr.",name + "!",
		"\nVersion:", version,
	)
	showMenu()
}

// EXIBE O MENU E RETORNA A OPÇAÕ ESCOLHIDA
func showMenu() {

	opcao := -1
	// ATENÇÃO - ESSE "FOR" EQUIVALE AO "DO... WHILE" EM GO
	for ok := true; ok; ok = opcao != 0 {
		fmt.Println(
			"\n=========================",
			"\n1- Iniciar Monitoramento",
			"\n2- Exibir os logs",
			"\n0 - Sair do programa",
			"\n=========================",
		)
		opcao = readComand(opcao)
		choseOption(opcao)
	}
}

func readComand(opcao int) int {
	fmt.Print("\nDigite uma opção: ")
	fmt.Scan(&opcao) // só vai aceitar o tipo da variável - retorna 0 
	return opcao
}

// SWITCH RESPONSÁVEL POR TRATAR A OPÇÃO ESCOLHIDA NO MENU
func choseOption(opcao int) {
	switch opcao {
		case 0:
			fmt.Println("\nFim do programa!")
			os.Exit(0)
		case 1:
			fmt.Println("\nMonitoramento")
		case 2:
			fmt.Println("\nLogs")
		default:
			fmt.Println("\nOpção inválida")	
	}
}
