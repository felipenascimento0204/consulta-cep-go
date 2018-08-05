package postmon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Objeto que representa o retorno da api postmon
type Endereco struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
}

// Consulta um cep na api postmon
func ConsultarCep(cep string) {

	fmt.Println("")
	fmt.Println("### Iniciando consulta ###")
	fmt.Println("")

	endpoint := fmt.Sprintf("http://api.postmon.com.br/v1/cep/%s", url.QueryEscape(cep))

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var endereco Endereco

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&endereco); err != nil {
		log.Println(err)
	}

	fmt.Println("Logradouro:", endereco.Logradouro)
	fmt.Println("Bairro:", endereco.Bairro)
	fmt.Println("Cidade:", endereco.Cidade)
	fmt.Println("Estado:", endereco.Estado)
	fmt.Println("")
	fmt.Println("### Consulta finalizada ###")
	fmt.Println("")
}
