package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("the secret number is ", secretNumber)

	fmt.Println("Please input your guess")

	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			continue
		}

		input = strings.TrimSuffix(input, "\r\n")
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value", err)
			continue
		}

		fmt.Println("You guess is", guess)

		if guess == secretNumber {
			fmt.Println("You are right")
		} else if guess > secretNumber {
			fmt.Println("You guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("You guess is smaller than the secret number. Please try again")
		}
	}

}
