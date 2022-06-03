
#HELLO WORLD EM GOLANG
	o package main é como a classe main em java,
	indica que esse é o arquivo principal na hora de rodar
	funções se escreve com func;
	primeira letra da função de um pacote externo é MAIÚSCULA
	rodar- go run nome.go
	criar executável - go build nome.go

#VARIAVEIS
	Se declara o tipo depois do nome da variavel
	No go as variáveis não vem com lixo, vem zerada (AEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE)
	Não deixa rodar com variáveis não usadas
	LISTA RETORNA STRING VAZIA NAS POSIÇÕES NÃO USADAS
	fmt - BIBLIOTECA DE PRINT

#INFERÊNCIA E TIPOS
	O go entende que se uma var receber "", ela é string
	numero inteiro ele entende como int
	no caso do float, como tem 2 tipos, ele considera sempre o maior se não tiver nada
	reflet - BIBLIOTECA PRA CHECAR TIPO
	DA PRA DECLARAR VARIAVEIS IGUAL AO PYTHON 0_o (sem o var na frente) - nome : valor

#FUNÇÕES
	Pra declarar os atributos, o tipo também vai depois do nome

#CONDIÇÕES
    IF SEM PARENTESES
    if comando == 1 {
        // codigo
    } else if comando == 2 {
        // outra coisa
    }
    SWITCH NÃO PRECISA DE BREAK
    switch opcao {
		case 0:
			fmt.Println("\nFim do programa!")
		case 1:
			fmt.Println("\nMonitoramento")
		case 2:
			fmt.Println("\nLogs")
		default:
			fmt.Println("\nOpção inválida")			
	}

#LAÇOS DE REPETIÇÃO
	https://medium.com/rafaeltardivo/desbravando-o-go-estruturas-de-repeti%C3%A7%C3%A3o-parte-4-b0023f468b16
	ATENÇÃO - ESSE FOR EQUIVALE AO WHILE EM GO
	for op != 0 {
		fmt.Println(
			"1- Iniciar Monitoramento",
			"\n2- Exibir os logs",
			"0 - Sair do programa",
		)
		fmt.Scanf("%d", &op) //scanf igualzinho ao c
	}
#MÓDULOS
 - Na pasta raiz, dar um go module init workspace;
 -Quando precisar importar, dar um: 
 	import nome "workspace/pasta"
 -Na hora de chamar o arquivo importado, usar o nome dado.NomeDaFuncIniciandoComLetraMaiscula

 #SLICE 
	basicamente oq eu tinha usado (SÓ Q ELE NÃO OBRIGA PASSAR O TAMANHO): nomes := []string {1,2,3,4,5}
	Quando é necessário colocar mais elementos do que sua capacidade atual, o slice dobra a capacidade.

