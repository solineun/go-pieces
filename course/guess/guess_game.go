package guess

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

func StartGame() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(1000) + 1
	fmt.Println("number got chosen")
	reader := bufio.NewReader(os.Stdin)

	var result string
	if guessNumber(target, reader) {
		result = "Right, "
	} else {
		result = "Wrong, "
	}
	fmt.Println(result, target)
}

func guessNumber(target int, reader *bufio.Reader) bool {
	var guess int	
	for i := 0; i < 10; i++ {
		fmt.Print("guess the number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = parser(input)

		if guess < target {
			fmt.Println("too small, tries left: ", 9 - i)
		} else if guess > target {
			fmt.Println("too big, tries left: ", 9 - i)
		} else {
			return true
		}
	}
	return false
}

func parser(input string) int {
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	return guess
}