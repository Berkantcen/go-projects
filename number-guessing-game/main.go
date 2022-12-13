package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	loopOn := true
	secret := getRandomNumber()
	lives := 3
	fmt.Println("Welcome to number guessing game.")
	fmt.Printf("Make a guess between 1-10 and you have %d lives\n", lives)
	for loopOn {
		guess := getUserInput()
		compareGuessAndSecret(guess, secret, &lives, &loopOn)
	}
}

// Create a random number between 1-10
func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	number := rand.Int()%10 + 1
	return number
}

// Take users input as guess
func getUserInput() int {
	fmt.Println("Please input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to get your input")
	} else {
		fmt.Println("You guessed:", input)
	}
	return input
}

// Compare guess and secret
func compareGuessAndSecret(guess int, secret int, lives *int, loopOn *bool) {
	if guess < secret {
		*lives--
		fmt.Println("Too low")
		if *lives == 0 {
			fmt.Println("You have no more lives")
			*loopOn = false
		} else {
			fmt.Printf("You have %d lives left\n", *lives)
		}

	} else if guess > secret {
		fmt.Println("Too high")
		*lives--
		if *lives == 0 {
			fmt.Println("You have no more lives")
			*loopOn = false
		} else {
			fmt.Printf("You have %d lives left\n", *lives)
		}
	} else {
		fmt.Println("You won!")
		*loopOn = false
	}
}
