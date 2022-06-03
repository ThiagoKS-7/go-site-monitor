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
			initMonitoring()
		case 2:
			initShowingLogs()
		default:
			fmt.Println("\nOpção inválida")	
	}
}

/*Faz uma requisição http pra cada site, conferindo se retorna 200*/
func initMonitoring() {
	fmt.Println("\nMonitorando...")
	for _, i := range sites {
		res, err := http.Get(i)
		res = handleRequest(res,err)
		doMonitoring(res) 
	}
}

func handleRequest(res *http.Response,err error) *http.Response {
	//Equivalência de try catch
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	return res
}

func doMonitoring(res *http.Response) {
	fmt.Println(res.Request.URL, res.Status)
}

/*Mostra, com dia, data e hora, os logs dos sites solicitados*/
func initShowingLogs() {
	fmt.Println("\nMostrando logs...")
	for _, i := range sites {
		res, err := http.Get(i)
		res = handleRequest(res,err)
		showLogs(res)
	}
}

func showLogs(res *http.Response) {
	fmt.Println(res.Request.URL, res.Status)
}
