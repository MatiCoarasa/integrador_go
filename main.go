package main

import (
	"fmt"
)

func main() {
	unEscenario := Escenario{nombre: "Escenario1", categoria: BajoCoste{zona: "Caballito"}}

	fmt.Println(unEscenario)
	unEscenario.evolucion(unEscenario.categoria)
	fmt.Println(famaDeCategoria(unEscenario.categoria))
	fmt.Println(unEscenario)
	fmt.Println(extras(unEscenario.categoria))

	unRecuerdo := Recuerdo{descripcion: "Viaje", escenario: unEscenario}

	tomas := Anfitrion{personaje: Personaje{energia: 3}, velocidadProcesamiento: 5, recuerdos: []Recuerdo{unRecuerdo}}
	fmt.Println(tomas)

	fmt.Println(felicidadDeUnPersonaje(tomas))
}
