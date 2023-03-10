package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

// Busca e retorna o cep para o client
func BuscarCep() *cli.App {
	app := cli.NewApp()
	app.Name = "Buscar cep"
	app.Usage = "Busca o cep informado e retorna a cidade"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "cep",
			Value: "55641715",
			Usage: "Busca o cep informado",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "buscar",
			Flags:  flags,
			Action: buscar,
		},
	}

	app.Run(os.Args)
	return app
}

func buscar(c *cli.Context) {
	var url string = "https://viacep.com.br/ws/" + c.String("cep") + "/json/"
	response, erro := http.Get(url)

	if erro != nil {
		log.Fatal("Ouve um erro ao procurar pelo cep")
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		log.Fatal("Ouve um erro ao ler a resposta")
	}

	// bodyString := string(body)
	defer fmt.Println(bodyString)
	defer fmt.Println("Mostrando os dados da cidade")
}
