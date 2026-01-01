package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			begin := time.Now()
			fmt.Println(r.Method)
			next.ServeHTTP(w, r)
			fmt.Println("Requisição " + r.Method + r.URL.String() + " processada em:", time.Since(begin))
		},
	)
}

func main() {	
	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /healthcheck",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method)
			fmt.Fprintln(w, "hello, world")
		},
	)

	mux.HandleFunc(
		"GET /user/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.PathValue("id"))
		},
	)


	srv := &http.Server{
		Addr: ":8080",
		Handler: log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil { // trava a execução do programa
		if !errors.Is(err, http.ErrServerClosed){
			panic(err)
		} 
	}
}
