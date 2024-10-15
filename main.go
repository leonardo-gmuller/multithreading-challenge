package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/leonardo-gmuller/multitheading-challenge/dto"
)

const cep string = "01153000"

func main() {
	ch1 := make(chan dto.BrasilApiResponse)
	ch2 := make(chan dto.ViaCepResponse)

	go SearchCepBrasilApi(ch1)
	go SearchCepViaCep(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println("Resposta da BrasilAPI:")
		printJson(msg1)

	case msg2 := <-ch2:
		fmt.Println("Resposta da ViaCEP:")
		printJson(msg2)

	case <-time.After(time.Second):
		println("timeout")
	}

}

func SearchCepViaCep(ch chan dto.ViaCepResponse) {
	const url = "http://viacep.com.br/ws/"
	req, err := http.Get(url + cep + "/json/")
	if err != nil {
		fmt.Println("Erro ao acessar ViaCEP:", err)
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Erro ao ler resposta da ViaCEP:", err)
		return
	}

	var data dto.ViaCepResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Erro ao fazer unmarshal da resposta da ViaCEP:", err)
		return
	}

	ch <- data
}

func SearchCepBrasilApi(ch chan dto.BrasilApiResponse) {
	const url = "https://brasilapi.com.br/api/cep/v1/"
	req, err := http.Get(url + cep)
	if err != nil {
		fmt.Println("Erro ao acessar BrasilAPI:", err)
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Erro ao ler resposta da BrasilAPI:", err)
		return
	}

	var data dto.BrasilApiResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Erro ao fazer unmarshal da resposta da BrasilAPI:", err)
		return
	}

	ch <- data
}

func printJson(data interface{}) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		fmt.Println("Erro ao codificar para JSON:", err)
		return
	}
	fmt.Printf("%s\n", buf.Bytes())
}
