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

	tomas := Anfitrion{Personaje: Personaje{energia: 3}, velocidadProcesamiento: 5, recuerdos: []Recuerdo{unRecuerdo}}
	fmt.Println(tomas)

	fmt.Println(felicidadDeUnPersonaje(&tomas))
	interaccion(&tomas)
	fmt.Println(tomas)

	unEscenario.consecuencias(&tomas)
	fmt.Println(tomas)
	conocerEscenario(&tomas, unEscenario)
	fmt.Println(tomas)

	fmt.Println("Una Trama")
	unaTrama := Trama{personajes: []IPersonaje{&tomas}, escenario: unEscenario}
	fmt.Println(unaTrama)
	unaTrama.renovar()
	fmt.Println(unaTrama)

	fmt.Println("Nicolas")
	nicolas := Huesped{Personaje: Personaje{energia: 15}, minutosRestantes: 10, amigos: []IPersonaje{&tomas}}
	fmt.Println(nicolas)

	fmt.Println(felicidadDeUnPersonaje(&nicolas))
	fmt.Println(esRebelde(&tomas))

	otroEscenario := Escenario{nombre: "EscenarioDeLujo", categoria: DeLujo{visitas: 7}}
	fmt.Println(otroEscenario)
	otroEscenario.evolucion(otroEscenario.categoria)
	fmt.Println(otroEscenario)

}
