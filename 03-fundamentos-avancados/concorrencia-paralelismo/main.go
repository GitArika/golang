// Package main -> Concorrência e Paralelismo
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	requisicaoAsync()
	requisicaoParalela()
}

func requisicaoAsync() {
	start := time.Now()
	for range 10 {
		resp, err := http.Get("https://google.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // Faz a limpeza da conexão
		fmt.Println("ok")
	}
	fmt.Println("tempo de processamento assíncrono:",time.Since(start))
}

func requisicaoParalela() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			resp, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close() // Faz a limpeza da conexão
			fmt.Println("ok")
		}()
	}

	wg.Wait()
	fmt.Println("tempo de processamento concorrente:",time.Since(start))
}

