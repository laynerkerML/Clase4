package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ejecución finalizada")
	}()
	datos, err := os.ReadFile("../customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	datosStr := string(datos)
	fmt.Println("hola", datosStr)
}
