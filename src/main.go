package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type calc struct{}

func (calc) operate(entry string) int {
	re := regexp.MustCompile(`\D`)
	values := re.Split(entry, 2)
	operator := string(re.Find([]byte(entry)))

	value1 := parse(values[0])
	value2 := parse(values[1])

	switch operator {
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "*":
		return value1 * value2
	case "/":
		return value1 / value2
	default:
		log.Fatalf("unsupported operator %s\n", operator)
		return 0
	}
}

func parse(value string) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("first operator must be a number")
	}
	return parsed
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	operation := readEntry()
	c := new(calc)
	result := c.operate(operation)
	fmt.Println(result)
}
