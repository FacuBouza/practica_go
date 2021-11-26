package main

import (
	"fmt"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Finaliza")
}

func main() {
	ini := time.Now()
	proceso(1)
	proceso(2)
	proceso(3)
	fin := time.Now()

	tiempoTot := fin.Sub(ini)
	fmt.Printf("El tiempo secuencial total es %v\n", tiempoTot)

	ini = time.Now()
	for i := 0; i < 12; i++ {
		go proceso(i)
	}
	time.Sleep(2000 * time.Millisecond)

	fin = time.Now()

	tiempoTot = fin.Sub(ini)
	fmt.Println("El tiempo paralelo total es ", tiempoTot.Seconds())
}
