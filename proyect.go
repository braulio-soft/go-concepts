package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	numRand := rand.Intn(100)
	var (
		numberEnter   int
		numberOptions int
	)
	const maxOptions = 10
	for numberOptions < maxOptions {
		numberOptions++
		fmt.Printf("Enter your number (intentos restantes : %d): ", maxOptions-numberOptions+1)
		fmt.Scanln(&numberEnter)
		if numberEnter == numRand {
			fmt.Println("Congratulations ! Your answer its correct you guessed the number")
			playAgain()
			return
		} else if numberEnter < numRand {
			fmt.Println("The number to guess is higher")
		} else if numberEnter > numRand {
			fmt.Println("The number to guess is less")
		}
	}
	fmt.Println("You have no more options . The number was : ", numRand)
	playAgain()

}

func playAgain() {
	var election string
	fmt.Print("Do you want to play again? (y/n): ")
	fmt.Scanln(&election)
	switch election {
	case "y":
		play()
	case "n":
		fmt.Println("Thanks for playing ")
	default:
		fmt.Println("Wrong answer! Try Again! ")
		playAgain()
	}
}
