package main

//Problema: el lenguaje permite utilizar interfaces para los métodos o composición para simular la herencia (el estado), lo que necesitamos es una combinación de ambos.

type Personaje struct {
	energia int
}

type Anfitrion struct {
	Personaje
	velocidadProcesamiento int
	recuerdos              []Recuerdo
}

type Huesped struct {
	Personaje
	minutosRestantes int
	amigos           []IPersonaje
}

type Parque struct {
	factor int
}

type IPersonaje interface {
	sumarEnergia(cantidad int)
	felicidad() int
	interactuar()
	consecuenciasDeConocerEscenario(e Escenario)
	interactuarConMuchos(personajes []IPersonaje)
}

func (h *Huesped) sumarEnergia(unaCantidad int) {
	(*h).energia = h.energia - unaCantidad
}

func (a *Anfitrion) sumarEnergia(unaCantidad int) {
	(*a).energia = a.energia - unaCantidad
}

func sumarEnergiaDeUnPersonaje(i IPersonaje, unaEnergia int) {
	i.sumarEnergia(unaEnergia)
}

func felicidadDeUnPersonaje(i IPersonaje) int {
	return i.felicidad()
}

func (h Huesped) felicidad() int {
	return h.minutosRestantes * h.felicidadDeAmigos()
}

func (h Huesped) felicidadDeAmigos() int {
	felicidadTotalDeAmigos := 0
	amigosDelPersonaje := h.amigos
	for _, IPersonaje := range amigosDelPersonaje {
		felicidadTotalDeAmigos += felicidadDeUnPersonaje(IPersonaje)
	}
	return felicidadTotalDeAmigos
}

func (a Anfitrion) felicidad() int {
	unParque := Parque{factor: 10}
	return a.energia / a.velocidadProcesamiento * unParque.factor
}

func (p *Personaje) interactuar() {
	(*p).energia = p.energia / 2
}

func interaccion(i IPersonaje) {
	i.interactuar()
}

func (a *Anfitrion) interactuar() {
	(*a).Personaje.interactuar()
	(*a).velocidadProcesamiento = a.velocidadProcesamiento / 2
}

func (h *Huesped) interactuar() {
	(*h).Personaje.interactuar()
	(*h).minutosRestantes = h.minutosRestantes - 10
}

func rebeldia(p IPersonaje) int {
	return 1 / felicidadDeUnPersonaje(p)
}

func esRebelde(p IPersonaje) bool {
	return rebeldia(p) > 10

}

func (a Anfitrion) interactuarConMuchos(personajes []IPersonaje) {
	for _, personaje := range personajes {
		personaje.interactuar()
	}
}

func (h Huesped) interactuarConMuchos(personajes []IPersonaje) {
	for _, personaje := range personajes {
		personaje.interactuar()
	}
}

func interaccionTotal(i IPersonaje, personajes []IPersonaje) {
	i.interactuarConMuchos(personajes)
}

func (a *Anfitrion) consecuenciasDeConocerEscenario(e Escenario) {
	(*a).recuerdos = append(a.recuerdos, Recuerdo{descripcion: "Conoci un escenario", escenario: e})
}

func (h *Huesped) consecuenciasDeConocerEscenario(e Escenario) {
	(*h).minutosRestantes = h.minutosRestantes - 10
}

func (e Escenario) consecuencias(i IPersonaje) {
	i.consecuenciasDeConocerEscenario(e)
}

func (p *Personaje) consecuenciasDeConocerEscenario(e Escenario) {}

func conocerEscenario(p IPersonaje, e Escenario) {
	sumarEnergiaDeUnPersonaje(p, -e.fama())
	p.consecuenciasDeConocerEscenario(e)
	e.aumentarVisitas()
}
