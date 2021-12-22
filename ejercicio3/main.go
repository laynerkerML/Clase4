package main

import (
	"fmt"
	"os"
)

func myCustomErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 1000

	result, err := myCustomErrorTest(salary)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
