package main

import (
	"buscarcep/app"
	"fmt"
)

func main() {
	defer recuperarExecucao()
	dadosCidade := app.BuscarCep()
	var mostrarNovamente int
	fmt.Println(dadosCidade)
	fmt.Print("\n Deseja fazer uma nova busca ? 1 para sim ou 2 para não: ")
	_, erro := fmt.Scan(&mostrarNovamente)

	if erro != nil {
		panic(erro)
	}

	if mostrarNovamente == 1 {
		main()
	}
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução Recuperada")
		main()
	}
}
