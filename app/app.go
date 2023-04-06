package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Buscar e retorna o cep para o client
func Buscar(cep string) (dadosCidade string, result int) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	response, erro := http.Get(url)

	if erro != nil {
		erro := "ouve um erro ao encontrar o cep"
		return erro, 403
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		erro := "ouve um erro ao ler a resposta"
		return erro, 403
	}

	bodyString := string(body)

	if strings.Contains(bodyString, "400") || strings.Contains(bodyString, "erro") {
		erro := "ouve um erro com a sua requisição"
		return erro, 403
	}

	return bodyString, 200
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
