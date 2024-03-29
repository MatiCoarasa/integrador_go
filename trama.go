package main

type Trama struct {
	personajes []IPersonaje
	escenario  Escenario
}

func (t Trama) felicidadDePersonajes() int {
	felicidadDePersonajes := 0
	for _, personaje := range t.personajes {
		felicidadDePersonajes += felicidadDeUnPersonaje(personaje)
	}
	return felicidadDePersonajes
}

func (t Trama) complejidad() int {
	return t.escenario.fama() / t.felicidadDePersonajes()
}

func (t Trama) renovar() {
	t.escenario.evolucion(t.escenario.categoria)
	t.filtrarPersonajesDeTrama()
}

func (t *Trama) filtrarPersonajesDeTrama() {
	var personajesRebeldes []IPersonaje

	(*t).personajes = personajesRebeldes
	for _, personaje := range t.personajes {
		if esRebelde(personaje) {
			personajesRebeldes = append(personajesRebeldes, personaje)
		}
	}
}

func (t Trama) cruzar(o Trama) {
	t.personajesConocenEscenario(o.escenario)
	t.cruzarPersonajes(o.personajes)
}

func (t Trama) personajesConocenEscenario(unEscenario Escenario) {
	for _, personaje := range t.personajes {
		conocerEscenario(personaje, unEscenario)
	}
}

func (t Trama) cruzarPersonajes(o []IPersonaje) {
	for _, personaje := range t.personajes {
		interaccionTotal(personaje, o)
	}
}

func (t Trama) esAburrida() bool {
	return t.complejidad() < 10
}
