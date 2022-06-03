package menu

import (
	"fmt"
	"net/http"
	"os"
)

// EXIBE O MENU E RETORNA A OPÇAÕ ESCOLHIDA
func ShowMenu() {

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
		opcao = ReadComand(opcao)
		ChoseOption(opcao)
	}
}

func ReadComand(opcao int) int {
	fmt.Print("\nDigite uma opção: ")
	fmt.Scan(&opcao) // só vai aceitar o tipo da variável - retorna 0 
	return opcao
}

// SWITCH RESPONSÁVEL POR TRATAR A OPÇÃO ESCOLHIDA NO MENU
func ChoseOption(opcao int) {
	switch opcao {
		case 0:
			fmt.Println("\nFim do programa!")
			os.Exit(0)
		case 1:
			InitMonitoring()
		case 2:
			ShowLogs()
		default:
			fmt.Println("\nOpção inválida")	
	}
}




/*Faz uma requisição http pra cada site, conferindo se retorna 200*/
func InitMonitoring() {
	fmt.Println("\nMonitorando...")
	site := "https://www.alura.com.br"
	response, error := http.Get(site)
	//Equivalência de try catch
	if error != nil {
		fmt.Println(response)
	} else {
		fmt.Println("ERROR:", error)
	}
	
}
/*Mostra, com dia, data e hora, os logs dos sites solicitados*/
func ShowLogs() {
	fmt.Println("\nLogs")
}
