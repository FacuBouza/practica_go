package main

import (
	"fmt"
	"time"
)

func proceso(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Finaliza")
	c <- i
}

func main() {
	c := make(chan int)

	ini := time.Now()
	for i := 0; i < 12; i++ {
		go proceso(i, c)
	}

	for i := 0; i < 12; i++ {
		variable := <-c
		fmt.Println(variable)
	}

	fin := time.Now()

	tiempo := fin.Sub(ini)
	fmt.Println("El tiempo total de ejecución es ", tiempo.Seconds())
}
