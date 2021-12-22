package main

import (
	"errors"
	"fmt"
)

func myCustomErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("el salario ingresado no alcanza el mínimo imponible")
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 1000

	result, err := myCustomErrorTest(salary)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(result)
}
