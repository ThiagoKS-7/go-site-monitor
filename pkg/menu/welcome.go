package menu

import "fmt"

/*
	GetNameAndVersion()
	* TRATA NOME E VERSÃO
	@void
	* retruns name string, version float32
*/
func GetNameAndVersion() (string, float32) {
	name := "Thiago" // inferência na string
	var version float32 = 1.2 // aqui é bom deixar, pq tem 2 tipos de float
	return name, version
}

/*
	Welcome()
	* MOSTRA O NOME, VERSÃO E CHAMA O MENU
	@param name string
	@param version float32
*/
func Welcome(name string, version float32) {
	fmt.Println(
		"Hello, Mr.",name + "!",
		"\nVersion:", version,
	)
	ShowMenu()
}
