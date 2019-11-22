package main

type Trama struct {
	personajes []Personaje
	escenario  Escenario
}

func (t Trama) felicidadDePersonajes() int {
	felicidadDePersonajes := 0
	for _, num := range t.personajes {
		felicidadDePersonajes += num.felicidad()
	}
	return felicidadDePersonajes
}

func (t Trama) complejidad() int {
	return t.escenario.fama() / t.felicidadDePersonajes()
}

func (t Trama) renovar() {
	t.escenario.evolucionar()
	//Filter
}

func (t Trama) cruzar(o Trama) {
	t.personajesConocenEscenario(o.escenario)
	t.cruzarPersonajes(o.personajes)
}

func (t Trama) personajesConocenEscenario(unEscenario Escenario) {
	for _, num := range t.personajes {
		num.conocerEscenario(unEscenario)
	}
}

func (t Trama) cruzarPersonajes(o []Personaje) {
	for _, num := range t.personajes {
		//interactuarConMuchos(num, o)
	}
}

func (t Trama) esAburrida() bool {
	return t.complejidad < 10
}
