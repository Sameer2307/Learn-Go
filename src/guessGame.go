// Guess a number with 10 chances

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Generate a random number
	// get the current time and date as integer to improvr randomness
	seconds := time.Now().Unix()
	// seed the randomness
	rand.Seed(seconds)
	// rand.Intn(100) generates a random number from 0 to 99 and hence we are doing +1
	target := rand.Intn(100) + 1
	fmt.Println("I have selected a number from 1-100, Can you guess the number in 10 chances")

	success := false

	//looping to allow only 10 guesses
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have 10 guesses, guesses left", 10-guesses)
		// reading the integer
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Converting String input to Int
		// Trim the new line at the end
		input = strings.TrimSpace(input)
		// convert it into integer using the strconv
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Your guess was HIGHER than the target")
		} else if guess < target {
			fmt.Println("Your guess was LOWER than the target")
		} else {
			// set success true so that we know its done
			success = true
			fmt.Println("Good Job you guessed it!!!")
			// this break will make sure we exit the loop so that once correct guess we dont ask for more guesses
			break
		}
	}
	// if success was false, return that you lost the game
	if !success {
		fmt.Println("Sorry you were unable to guess the number, it was:", target)
	}
}
