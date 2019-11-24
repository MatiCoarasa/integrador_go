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

	tomas := Anfitrion{personaje: Personaje{energia: 20}, velocidadProcesamiento: 5, recuerdos: []Recuerdo{unRecuerdo}}
	fmt.Println(tomas)

	fmt.Println(felicidadDeUnPersonaje(&tomas))
	interaccion(&tomas)
	fmt.Println(tomas)

	unEscenario.consecuencias(&tomas)
	fmt.Println(tomas)
	tomas.personaje.conocerEscenario(unEscenario)
	fmt.Println(tomas)

	unaTrama := Trama{personajes: []Personaje{tomas.personaje}, escenario: unEscenario}
	fmt.Println(unaTrama)
	unaTrama.renovar()
	fmt.Println(unaTrama)

	nicolas := Huesped{personaje: Personaje{energia: 15}, minutosRestantes: 10, amigos: []Personaje{tomas.personaje, tomas.personaje}}
	fmt.Println(nicolas)

	fmt.Println(felicidadDeUnPersonaje(&nicolas))
	fmt.Println(tomas.personaje.esRebelde())
}
