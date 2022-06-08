package menu

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

/*	ShowMenu()
	*EXIBE O MENU E RETORNA A OPÇÃO ESCOLHIDA
	@void
*/
func ShowMenu() {

	opcao := -1
	// ?ATENÇÃO - ESSE "FOR" EQUIVALE AO "DO... WHILE" EM GO
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

//-----------------------END MENU -----------------------------------
/*
	* MAIN AUX'S *
*/


/*	ReadComand()
	* TRATA A OPÇÃO DIGITADA
	@param opcao: int
	* returns: valor da opção digitada
*/
func ReadComand(opcao int) int {
	fmt.Print("\nDigite uma opção: ")
	fmt.Scan(&opcao) // só vai aceitar o tipo da variável - retorna 0 
	return opcao
}

/*	ChoseOption()
	* TRATA OS CASES DO MENU
	@param opcao: int
	* returns: case da opção digitada
*/
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

/* 	doMonitoring()
	* TRATA O PRINT DO MONITORAMENTO
	@param res: *http.Response
	* returns display do monitoramento
*/ 
func doMonitoring(res *http.Response, site string) {
	registerLog(site, res.Status)
	fmt.Println(res.Request.URL, res.Status)
}

/* 	showLogs()
	* TRATA O PRINT DA URL LOGADA
	@param res: *http.Response
	* returns display dos logs
*/ 
func showLogs(res *http.Response) {
	fmt.Println("[" + time.Now().Format("02/01/2006 15:04:05") + "]"  + " - ",res.Request.URL, res.Status)
}

//------------------- END MAIN AUX'S--------------------------------------
/*
	* MAIN CASES *
*/


/* 	initMonitoring()
	*FAZ UMA REQUISIÇÃO HTTP PARA CADA SITE
	*JOGANDO A RESP E O POSSIVEL ERRO NA handleRequest().
	*SE TIVER RES, FAZ O MONITORAMENTO
	@void
*/ 
func initMonitoring() {
	fmt.Println("\nMonitorando...")
	fmt.Println("")
	files := readFileLinks()
	//Se retornar um arquivo, sites recebe ele,
	//senão, usa como padrão os descritos em sites.go
	if len(files) > 0 {
		sites = files
	}

	for i := 0; i < monitoramentos; i++ {
		fmt.Print("-----------------------TESTE #")
		fmt.Print(i + 1)
		fmt.Println("-----------------------")
		for _, i := range sites {
			res, err := http.Get(i)
			res = handleRequest(res,err)
			doMonitoring(res, i)
		}
		fmt.Println("")
		registerLog("\n","\n")
	}
	fmt.Println("\nFim do monitoramento!")
}

/*	initShowingLogs()
	* REQUISITA OS SITES E MANDA O RESULTADO PRA showLogs()
	@void
*/ 
func initShowingLogs() {
	fmt.Println("\nMostrando logs...")
	fmt.Println("")
	files := readFileLinks()
	//Se retornar um arquivo, sites recebe ele,
	//senão, usa como padrão os descritos em sites.go
	if len(files) > 0 {
		sites = files
	}
	for i := 0; i < monitoramentos; i++ {
		fmt.Print("-----------------------TESTE #")
		fmt.Print(i + 1)
		fmt.Println("-----------------------")
		for _, item := range sites {
			res, err := http.Get(item)
			res = handleRequest(res,err)
			showLogs(res)
		}
		fmt.Println("")
	}
	fmt.Println("\nFim dos logs!")
}

//-----------------------END MAIN CASES-------------------------------
/*
	* VALIDATIONS *
*/


/*  
	handleRequest()
	*VALIDA SE A ROTA NÃO RETORNOU ERRO
	@param res: *http.Response
	@param err: error
	*returns: *http.Response
*/ 
func handleRequest(res *http.Response,err error) *http.Response {
	// ?Equivalência de try catch em outras linguagens
	if err != nil {
		fmt.Println("Erro:", err)
	}
	defer res.Body.Close()
	return res
}

//---------------------------END VALIDATIONS------------------------------
/*
	* FILE MANAGERS *
*/

/* 
	readFileLinks()
	*ABRE O ARQUIVO SITES.TXT E GUARDA CADA LINHA NUMA POSIÇAO DO ARRAY SITES
	@void
	* returns []string
*/
func readFileLinks() []string {
	var sites []string
	// ?file, err := ioutil.ReadFile("../sites.txt") -- serve pra conseguir dar display de tudo
	file, err := os.Open("../sites.txt")
	if err != nil {
		fmt.Println("Arquivo não encontrado: ", err)
		fmt.Println("Usando sites default ao invés disso...")
		fmt.Println("")
	} else {
		reader := bufio.NewReader(file)
		//Vai lendo a linha até achar o EOF
		for {
			line, err := reader.ReadString('\n') // byte usa aspa simples
			line = strings.TrimSpace(line)
			sites = append(sites, line)
			if err == io.EOF {
				break
			}
		}
	}
	file.Close()
	return sites
}

/* 
	registerLog()
	*GUARDA URL, HORA E STATUS DE CADA SITE NUMA DAS LINHAS DE UM .TXT
	@param res *http.Response
*/
func registerLog(url string, status string) {
	file, err := os.OpenFile("../logs.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	if (url != "\n" && status != "\n") {
		file.WriteString( "[" + time.Now().Format("02/01/2006 15:04:05") + "]"  + " - " + url + " - status: " + status + "\n")
	} else {
		file.WriteString("\n")
	}
	file.Close();
}