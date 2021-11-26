package main

import "fmt"

func main() {
	var p1 *int
	var num int

	fmt.Printf("Antes de inicializar: %v, %v\n", p1, num)

	p1 = &num
	num = 5
	fmt.Printf("Despues de inicializar: %v, %v\n", p1, *p1)
	fmt.Printf("Despues de inicializar: %v, %v\n", &num, num)

	*p1 = 30
	fmt.Printf("Despues de inicializar: %v, %v\n", num, *p1)

	num2 := 25.6
	p2 := &num2

	fmt.Printf("Valor 2:", num, *p2)
}
