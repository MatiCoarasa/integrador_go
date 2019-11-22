package main

type Trama struct {
	personajes []Personaje
	escenario  Escenario
}

func (t Trama) felicidadDePersonajes() int {
	felicidadDePersonajes := 0
	for _, personaje := range t.personajes {
		felicidadDePersonajes += personaje.felicidad()
	}
	return felicidadDePersonajes
}

func (t Trama) complejidad() int {
	return t.escenario.fama() / t.felicidadDePersonajes()
}

func (t Trama) renovar() {
	t.escenario.evolucionar()
	t.filtrarPersonajesDeTrama()
}

func (t *Trama) filtrarPersonajesDeTrama() {
	var personajesRebeldes []Personaje
	(*t).personajes = personajesRebeldes
	for _, personaje := range t.personajes {
		if personaje.esRebelde() {
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
		personaje.conocerEscenario(unEscenario)
	}
}

func (t Trama) cruzarPersonajes(o []Personaje) {
	for _, personaje := range t.personajes {
		personaje.interactuarConMuchos(o)
	}
}

func (t Trama) esAburrida() bool {
	return t.complejidad() < 10
}
