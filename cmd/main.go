package main

import (
	run "workspace/pkg/menu"
)

/*
	* PROJETO DE MONITORAMENTO DE SISTEMAS EM GOLANG
	Exemplo da aula introdutória de go
	Consegue:
		-Iniciar monitoramento dos sites;
		-Exibir logs dos sites;
*/
func main() {
	run.Welcome(run.GetNameAndVersion())
}

