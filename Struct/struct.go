package main

import (
	"fmt"
	"reflect"
)

type Fecha struct {
	dia, mes, anio int
}

type Persona struct {
	nombre    string
	apellido  string
	direccion string
	edad      int
	fechaNac  Fecha
}

func main() {
	p1 := Persona{}
	fmt.Printf("%v\n%+v\n%T\n", p1, p1, p1)
	p1.nombre = "Matias"
	fmt.Printf("%v\n%+v\n%T\n", p1, p1, p1)

	p2 := Persona{"Facundo", "Bouza", "Argentina", 21, Fecha{1, 2, 2021}}
	fmt.Printf("%v\n%+v\n%T\n", p2, p2, p2)

	p3 := Persona{
		nombre:   "Facundo",
		fechaNac: Fecha{21, 12, 1989},
	}
	//Los campos que dejo sin inicializar quedan en 0 o en vacío si son string. No son nulos.
	fmt.Printf("%v\n%+v\n%T\n", p3, p3, p3)

	// No se puede inicializar en nil
	// p4 := Persona{
	// 	nombre: nil,
	// }

	tipo := reflect.TypeOf(p1)
	fmt.Println("Reflect")

	fmt.Println(tipo.Field(0))
	fmt.Println(tipo.Name())
	fmt.Println(tipo.Kind())

	// valor := reflect.ValueOf(p1)
	// fmt.Println("Reflect 2")

}
