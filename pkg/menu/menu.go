package menu

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

/* EXIBE O MENU E RETORNA A OPÇÃO ESCOLHIDA */
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

/* RETORNA A OPÇÃO DIGITADA PELO USUÁRIO */
func ReadComand(opcao int) int {
	fmt.Print("\nDigite uma opção: ")
	fmt.Scan(&opcao) // só vai aceitar o tipo da variável - retorna 0 
	return opcao
}

/* SWITCH RESPONSÁVEL POR TRATAR A OPÇÃO ESCOLHIDA NO MENU */
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

/*  FAZ UMA REQUISIÇÃO HTTP PARA CADA SITE
	JOGANDO A RESP E O POSSIVEL ERRO NA handleRequest.
	SE TIVER RES, FAZ O MONITORAMENTO
*/ 
func initMonitoring() {
	fmt.Println("\nMonitorando...")
	for i := 0; i < monitoramentos; i++ {
		for _, i := range sites {
			res, err := http.Get(i)
			res = handleRequest(res,err)
			doMonitoring(res) 
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

/*  VALIDA SE A ROTA NÃO RETORNOU ERRO */ 
func handleRequest(res *http.Response,err error) *http.Response {
	//Equivalência de try catch
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	return res
}

/*  MOSTRA A URL E O STATUS DE CADA SITE MONITORADO */ 
func doMonitoring(res *http.Response) {
	fmt.Println(res.Request.URL, res.Status)
}

/*  FAZ UMA REQUISIÇÃO HTTP PARA CADA SITE
	JOGANDO A RESP E O POSSIVEL ERRO NA handleRequest.
	SE TIVER RES, MOSTRA OS LOGS
*/ 
func initShowingLogs() {
	fmt.Println("\nMostrando logs...")
	for i := 0; i < monitoramentos; i++ {
		for _, i := range sites {
			res, err := http.Get(i)
			res = handleRequest(res,err)
			showLogs(res)
		}
	}
}

/*  MOSTRA A URL E O STATUS DE CADA SITE MONITORADO */ 
func showLogs(res *http.Response) {
	fmt.Println(res.Request.URL, res.Status)
}
