# Primeiros passos

Inicialize o seu modulo com go init:

```sh
go mod init my-first-go-project
```

## Escrevendo seu hello world

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Executando seu hello world

```sh
go run main.go
```

## Compilando seu hello world

```sh
go build main.go
```

Para alterar o arquivo gerado no build você pode usar a flag -o

```sh
go build -o "foo.exe" main.go
```

## Executando seu hello world

```sh
./main
```

## Compilando para windows

Utilizando a variavel de ambiente `GOOS` e `GOARCH`, podemos modificar o compilador go

```sh
GOOS=windows GOARCH=amd64 go build main.go
```

## Compilando para linux

```sh
GOOS=linux GOARCH=amd64 go build main.go
```

## Compilando para ARM

```sh
GOOS=darwin GOARCH=arm64 go build main.go
```

More info see [GOOS and GOARCH](./docs/goos-goarch.md)