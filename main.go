package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
    var num1, num2 float64
	var operator string

	fmt.Println("Masukkan angka pertama: ")
	fmt.Scanln(&num1)
	fmt.Println("Masukkan angka kedua: ")
	fmt.Scanln(&num2)
	fmt.Println("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
			return
		}
	default:
		fmt.Println("Operator tidak valid")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}