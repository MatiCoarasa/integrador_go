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
	fmt.Println(unRecuerdo.emotividad())

	tomas := Anfitrion{Personaje: Personaje{3}, velocidadProcesamiento: 5, recuerdos: []Recuerdo{unRecuerdo}}
	fmt.Println(tomas)

	fmt.Println(felicidadDeUnPersonaje(&tomas))
	interaccion(&tomas)
	fmt.Println(tomas)

	unEscenario.consecuencias(&tomas)
	fmt.Println(tomas)
	tomas.conocerEscenario(unEscenario)
	fmt.Println(tomas)

	unaTrama := Trama{personajes: []Personaje{tomas.Personaje}, escenario: unEscenario}
	fmt.Println(unaTrama)
	unaTrama.renovar()
	fmt.Println(unaTrama)

	nicolas := Huesped{Personaje: Personaje{15}, minutosRestantes: 10, amigos: []Personaje{tomas.Personaje, tomas.Personaje}}
	fmt.Println(nicolas)

	fmt.Println(felicidadDeUnPersonaje(&nicolas))
	fmt.Println(tomas.Personaje.esRebelde())
}
