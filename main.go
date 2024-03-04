package main

import (
    "fmt"
    "bufio"
	"os"
	"strconv"

	"./calculator"
)

func main() {
    fmt.Println("Hello, World!")
    reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')

	fmt.Print("Masukkan angka pertama: ")
	num1Str, _ := reader.ReadString('\n')
	num1, _ := strconv.ParseFloat(num1Str[:len(num1Str)-1], 64)

	fmt.Print("Masukkan angka kedua: ")
	num2Str, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseFloat(num2Str[:len(num2Str)-1], 64)

	result := 0.0

	switch operator[:len(operator)-1] {
	case "+":
		result = calculator.Add(num1, num2)
	case "-":
		result = calculator.Subtract(num1, num2)
	case "*":
		result = calculator.Multiply(num1, num2)
	case "/":
		res, err := calculator.Divide(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		result = res
	default:
		fmt.Println("Operator tidak valid")
		return
	}

	fmt.Printf("Hasil: %v\n", result)
}