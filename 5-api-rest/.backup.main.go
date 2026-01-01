package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const googleURL = "https://www.google.com"

func main() {
	time.Sleep(time.Second)
	res, err := http.Get(googleURL)
	if err != nil {
		fmt.Println("não foi possível alcançar o url: ")
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ocorreu um erro ao ler a resposta")
		return
	}

	fmt.Println(string(data))

	requestWithContext()
}

func request() {
	req, err := http.NewRequest(http.MethodGet, googleURL, nil)
	if err != nil {
		fmt.Println("não foi possível criar a requisição")
		return
	}

	req.Header.Set("authorization", "Bearer token")
	req.Header.Set("accept", "application/json")
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("não foi possível alcançar o url: ")
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ocorreu um erro ao ler a resposta")
		return
	}

	fmt.Println(string(data))
}

func requestWithContext() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/healthcheck", nil)
	if err != nil {
		fmt.Println("não foi possível criar a requisição")
		return
	}

	req.Header.Set("authorization", "Bearer token")
	req.Header.Set("accept", "application/json")
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("não foi possível alcançar o url: ")
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ocorreu um erro ao ler a resposta")
		return
	}

	fmt.Println(string(data))
}
