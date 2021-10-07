package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operation := scanner.Text()

	values := strings.Split(operation, "+")

	operator1, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("First operator must be a number")
	}

	operator2, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal("Second operator must be a number")
	}

	fmt.Println(operator1 + operator2)
}
