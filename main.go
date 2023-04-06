package main

import (
	"buscarcep/app"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	defer recuperarExecucao()

	http.HandleFunc("/cep", func(resWriter http.ResponseWriter, req *http.Request) {
		cep := req.URL.Query().Get("cep")
		dadosCidade, result := app.Buscar(cep)
		resWriter.WriteHeader(result)
		io.WriteString(resWriter, dadosCidade)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução Recuperada")
		main()
	}
}
