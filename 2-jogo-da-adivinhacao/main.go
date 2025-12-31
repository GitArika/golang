package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100")

	x := rand.Int64N(101) // limite não inclusivo

	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("Qual é o seu palpite?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)
	
		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("O seu palpite precisa ser um número inteiro")
			return
		}

		switch {
			case guessInt < x:
				fmt.Println("Você errou. O número sorteado é maior que", guessInt)
			case guessInt > x:
				fmt.Println("Você errou. O número sorteado é menor que", guessInt)
			case guessInt == x:
				fmt.Printf("Parabéns, você acertou! O número era: %d. Você acertou em %d tentativas\n", x, i+1)
				return
		}

		guesses[i] = guessInt
	}

	fmt.Printf("Infelizmente, você não acertou o número, que era: %d. Você teve 10 tentativas\n", x)
}
