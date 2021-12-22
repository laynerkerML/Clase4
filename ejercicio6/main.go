package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellidos string
	DNI       int
	NumTlf    string
	Domicilio string
}

type Empresa struct {
	Clientes []Cliente
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ejecución finalizada")
	}()
	empresa := new(Empresa)
	cliente := Cliente{
		Nombre:    "Laynerker",
		Apellidos: "Guerrero",
		DNI:       95950356,
		NumTlf:    "1123513062",
		Domicilio: "Puan 415",
	}
	varlidarCampos, err := empresa.validarDatos(cliente)
	if err != nil {
		panic(err)
	}
	validateExistCliente, err := empresa.verificarCliente(cliente)
	if err != nil {
		panic(err)
	}
	if validateExistCliente && varlidarCampos {
		empresa.AddClient(cliente)
	}

	fmt.Println("Empresa: ", empresa)
}

func (e *Empresa) AddClient(cliente Cliente) {
	rand.Seed(time.Now().UnixNano())
	cliente.Legajo = rand.Intn(1000)
	e.Clientes = append(e.Clientes, cliente)
}

func (e *Empresa) verificarCliente(cliente Cliente) (bool, error) {

	datos, err := os.ReadFile("../customers.txt")
	if err != nil {
		return false, errors.New("el archivo indicado no fue encontrado o está dañado")
	}
	datosStr := string(datos)
	if datosStr == cliente.Nombre {
		return false, nil
	}
	return true, nil
}

func (e *Empresa) validarDatos(cliente Cliente) (bool, error) {
	v := reflect.ValueOf(cliente)
	typeOfS := v.Type()
	campos := []string{}

	for i := 0; i < v.NumField(); i++ {
		if typeOfS.Field(i).Name != "Legajo" {
			if v.Field(i).Interface() == "" || v.Field(i).Interface() == 0 || v.Field(i).Interface() == nil {
				campos = append(campos, typeOfS.Field(i).Name)
			}
		}
		//fmt.Printf("field: %s \t value:%v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
	if len(campos) != 0 {
		return false, errors.New(fmt.Sprintf("Existen campos en blanco %s", campos))
	}
	return true, nil
}
