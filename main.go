package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hola Mundo")

	unEscenario := Escenario{nombre: "Escenario1", categoria: BajoCoste{zona: "Caballito"}}
	unEscenario.evolucion(unEscenario.categoria)
	fmt.Println(unEscenario.fama())
}
