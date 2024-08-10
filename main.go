package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Jogo da adivinhação")
	fmt.Println("Um número será sorteado entre 0 e 100 e você terá 10 tentativas para acerta-lo")

	x := rand.Int63n(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chute)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chute)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns você acertou! O número era: %d\n"+
					"Você acertou na %da tentativa.\n"+
					"Essas foram as suas tentativas: %v\n", x, i+1, chutes[:i])
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Infelizmente, voê não acertou o número, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram as suas tentativas: %v", x, chutes)

}
