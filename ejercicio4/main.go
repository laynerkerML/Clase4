package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	caculo, err := calSalary(10, 180000.00)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("hola", fmt.Sprintf("%.2f", caculo))
}

func calSalary(horas int, cantidad float64) (float64, error) {
	if horas < 80 || horas <= -1 {
		return 0.0, errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salarioBase := float64(horas) * cantidad
	if salarioBase > 150000 {
		salarioBase -= salarioBase * 0.1
	}

	return salarioBase, nil
}

func calAguinaldos(maxSalary float64, cantidadMeses float64) (float64, error) {
	calculo := maxSalary / 12 * cantidadMeses

	return calculo, nil
}
