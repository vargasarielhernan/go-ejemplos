package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What you want to do: add, substract, multiply or divide")
	input, _ := reader.ReadString('\n')
	fmt.Println("What is your first namber")
	fNum, _ := reader.ReadString('\n')
	floatfNum, err := strconv.ParseFloat(fNum[:len(fNum)-2], 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("What is your second namber")
	sNum, _ := reader.ReadString('\n')
	floatsNum, err := strconv.ParseFloat(sNum[:len(sNum)-2], 64)
	if err != nil {
		panic(err)
	}
	switch input[:len(input)-2] {
	case "add":
		fmt.Println("Result:")
		fmt.Println(floatfNum + floatsNum)
	case "substract":
		fmt.Println("Result:")
		fmt.Println(floatfNum - floatsNum)
	case "multiply":
		fmt.Println("Result:")
		fmt.Println(floatfNum * floatsNum)
	case "divide":
		fmt.Println("Result:")
		fmt.Println(floatfNum / floatsNum)
	default:
		fmt.Println("invalid input")
	}

}
