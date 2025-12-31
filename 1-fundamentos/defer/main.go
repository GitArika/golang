package main

import (
	"fmt"
	"os"
	"database/sql"
)

func main() {
	doDefer()
	doMultipleDefer()
	doDeferFunc()
}

func doDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func doMultipleDefer() {
	defer fmt.Println(3)
	defer fmt.Println(2)

	fmt.Println(1)
}

func doDeferFunc() {
	x := 10

	defer func(y int){
		fmt.Println(y)
	}(x)

	fmt.Println("inside log", x)
	x = 50

	defer func(y int){
		fmt.Println(y)
	}(x)
}

func openFile() {
	file, err := os.Open("foo.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
}

var arquivos = [3]string{"foo.txt", "bar.txt", "baz.txt"}

func openMultipleFiles() {
	for _, f := range arquivos {
		func(){
			file, err := os.Open(f)
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}()
	}
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("localhost", "")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Causa um erro, pois fecha o DB antes do retorno do db

	return db, nil
}
