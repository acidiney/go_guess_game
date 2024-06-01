package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func game(magicNumber int) {
	var guess int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess a number between 1 and 100")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	guess, err = strconv.Atoi(strings.TrimSpace(text))

	if err != nil {
		fmt.Println(err)
	}

	if guess > magicNumber {
		fmt.Println("Your guess is too high")
		game(magicNumber)
	} else if guess < magicNumber {
		fmt.Println("Your guess is too low")
		game(magicNumber)
	} else {
		fmt.Println("Your guess is correct")
	}
}

func main() {
	fmt.Println("Hello World!")

	magicNumber := rand.IntN(100)

	game(magicNumber)
}
