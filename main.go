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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("lucky number: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func getNumberOfDieAndSides(text string) int {
	s := strings.Split(text, "d")
	numberOfDice, _ := strconv.Atoi(s[0])
	numberOfSides, _ := strconv.Atoi(s[1])
	total := 0
	for i := 1; i <= numberOfDice; i++ {
		total += randomise(numberOfSides)
	}
	return total
}

func randomise(num int) int {
	return rand.Intn(num-1) + 1
}
