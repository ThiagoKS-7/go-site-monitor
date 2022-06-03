package menu

import "fmt"

// EXEMPLO DE RETORNO DUPLO - RETORNA NOME E VERSÃO
func GetNameAndVersion() (string, float32) {
	name := "Thiago" // inferência na string
	var version float32 = 0.3 // aqui é bom deixar, pq tem 2 tipos de float
	return name, version
}

// MOSTRA NOME E CHAMA O MENU
func Welcome(name string, version float32) {
	fmt.Println(
		"Hello, Mr.",name + "!",
		"\nVersion:", version,
	)
	ShowMenu()
}
