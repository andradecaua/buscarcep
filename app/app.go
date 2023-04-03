package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Busca e retorna o cep para o client
func BuscarCep() string {
	defer recuperarExecucao()
	var cep string

	fmt.Println("------------------------")
	fmt.Println("Buscar Cep Iniciando")
	fmt.Print("\nPor gentileza informe o cep: ")
	_, err := fmt.Scan(&cep)

	if err != nil {
		panic(err)
	}
	return buscar(cep)
}

func buscar(cep string) (dadosCidade string) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	response, erro := http.Get(url)

	if erro != nil {
		erro := "ouve um erro ao encontrar o cep"
		return erro
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		erro := "ouve um erro ao ler a resposta"
		return erro
	}

	bodyString := string(body)

	if strings.Contains(bodyString, "400") || strings.Contains(bodyString, "erro") {
		erro := "ouve um erro com a sua requisição"
		return erro
	}

	defer fmt.Println("Mostrando os dados da cidade")
	return bodyString
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
