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
	operator := "-"

	values := strings.Split(operation, operator)

	value1, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("first operator must be a number")
	}

	value2, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal("second operator must be a number")
	}

	switch operator {
	case "+":
		fmt.Println(value1 + value2)
	case "-":
		fmt.Println(value1 - value2)
	case "*":
		fmt.Println(value1 * value2)
	case "/":
		fmt.Println(value1 / value2)
	default:
		log.Fatal("unsupported operator")
	}

}
