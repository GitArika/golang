package main

import (
	"errors"
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

var ErrNotFound = errors.New("not found")

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"não existe raiz quadrada de numero negativo"}
	}
	resultado := math.Sqrt(x)
	if math.Trunc(resultado) != resultado {
		return 0, ErrNotFound
	}

	return resultado, nil
}

func main() {
	a, b := 10, 3

	res, err := dividir(a,b)
	if err != nil { // tratamento de erro como valor
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	user, err := NewUser(false)
	if err != nil { // tratamento de erro como valor
		fmt.Println(err)
		return
	}
	user.Foo()
	var x float64 = 25

	resultado, err := raizQuadrada(x)
	if err != nil && errors.Is(err, ErrNotFound) { // validação erro customizado
		fmt.Println(err)
		return
	}

	var sqrtError SqrtError

	if err != nil && errors.As(err, &sqrtError) { // validação de type error 
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", resultado)

	errQualquer := baz()
	if errQualquer != nil && errors.Is(errQualquer, ErroQualquer) {
		fmt.Println("Deu erro em baz:", errQualquer)
	}

	erroJoin := c()

	fmt.Println(err)
	fmt.Println(errors.Is(erroJoin, ErroQualquer))
	fmt.Println(errors.Is(erroJoin, ErroQualquer2))
}

// interface pública disponível nativamente em Go
type error interface {
	Error() string
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não é possível fazer divisões por zero")
	}

	return a / b, nil
}

type User struct{
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("ocorreu um erro ao criar usuário")
	}

	return &User{"Fooooo"}, nil
}


func baz() error {
	err := tez()
	if err != nil {
		return fmt.Errorf("deu erro em tez: %w", err)
	}
	return nil
}

var ErroQualquer = errors.New("error")

func tez() error {
	return ErroQualquer
}

var ErroQualquer2 = errors.New("error 2")

func a() error { return ErroQualquer }
func b() error { return ErroQualquer2 }

func c() error {
	var errorResult error

	if err:= a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}
	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	return errorResult
}
