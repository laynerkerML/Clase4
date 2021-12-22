package main

import "fmt"

// sólo se requiere crear un tipo que implemente el método Error()
type myCustomError struct {
	msg string
}

//hacemos que nuestro tipo struct implemente el método Error()
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func myCustomErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", &myCustomError{
			msg: "algo salió mal",
		}
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 180000

	result, err := myCustomErrorTest(salary)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
