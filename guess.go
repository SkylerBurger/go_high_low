package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateTarget() (target int64) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	target = int64(rand.Intn(100) + 1)
	return
}

func solicitGuess(i int) (guess int64) {
	remaining := 10 - i
	fmt.Printf("Remaining Guesses: %d\n", remaining)
	fmt.Print("Guess a number between 1 and 100: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	guess, err := strconv.ParseInt(response, 10, 64)
	if err != nil {
		log.Fatal("Please enter a number.")
	}
	return
}

func checkGuess(target int64, guess int64) (gameOver bool) {
	if guess > target {
		fmt.Println("Your guess was too high.")
	} else if guess < target {
		fmt.Println("Your guess was too low.")
	} else {
		fmt.Println("You got it!!")
		gameOver = true
	}
	return
}

func main() {
	winner := false
	target := generateTarget()
	for i := 0; i < 10; i++ {
		guess := solicitGuess(i)
		gameOver := checkGuess(target, guess)
		if gameOver {
			winner = true
			break
		}
	}
	if !winner {
		fmt.Printf("Sorry you didn't win, the target was %d\n", target)
	}
	fmt.Println("Play again soon!")
}