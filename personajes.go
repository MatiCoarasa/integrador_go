package main

type Personaje struct {
	energia int
}

type Anfitrion struct {
	personaje              Personaje
	velocidadProcesamiento int
	recuerdos              []Recuerdo
}

type Huesped struct {
	personaje        Personaje
	minutosRestantes int
	amigos           []Personaje
}

type Parque struct {
	factor int
}

type IPersonaje interface {
	felicidad() int
	interactuar()
	consecuenciasDeConocerEscenario(e Escenario)
}

func felicidadDeUnPersonaje(i IPersonaje) int { // // implementacion de interfaz
	return i.felicidad()
}

func (p Personaje) felicidad() int {
	return p.felicidad()
}

func (h Huesped) felicidad() int {
	return h.minutosRestantes * h.felicidadDeAmigos()
}

func (h Huesped) felicidadDeAmigos() int {
	felicidadTotalDeAmigos := 0
	amigues := h.amigos
	for _, personaje := range amigues {
		felicidadTotalDeAmigos += personaje.felicidad()
	}
	return felicidadTotalDeAmigos
}

func (a Anfitrion) felicidad() int {
	unParque := Parque{factor: 10}
	return a.personaje.energia / a.velocidadProcesamiento * unParque.factor
}

func (p *Personaje) interactuar() {
	(*p).energia = p.energia / 2
}

func interaccion(i IPersonaje) { // // implementacion de interfaz
	i.interactuar()
}

func (a *Anfitrion) interactuar() {
	(*a).personaje.energia = a.personaje.energia / 2
	(*a).velocidadProcesamiento = a.velocidadProcesamiento / 2
}

func (h *Huesped) interactuar() {
	(*h).personaje.energia = h.personaje.energia / 2
	(*h).minutosRestantes = h.minutosRestantes - 10
}

func (p Personaje) rebeldia() int {
	return 1 / p.felicidad()
}

func (p Personaje) esRebelde() bool {
	return p.rebeldia() > 10

}

func (p Personaje) interactuarConMuchos(personajes []Personaje) {
	for _, personaje := range personajes {
		personaje.interactuar()
	}
}

func (a *Anfitrion) consecuenciasDeConocerEscenario(e Escenario) {
	(*a).recuerdos = append(a.recuerdos, Recuerdo{descripcion: "Conoci un escenario", escenario: e})
}

func (h *Huesped) consecuenciasDeConocerEscenario(e Escenario) {
	(*h).minutosRestantes = h.minutosRestantes - 10
}

func consecuencias(i IPersonaje, e Escenario) { // implementacion de interfaz
	i.consecuenciasDeConocerEscenario(e)
}

func (p *Personaje) consecuenciasDeConocerEscenario(e Escenario) {}

func (p *Personaje) conocerEscenario(e Escenario) {
	(*p).energia = p.energia - e.fama()
	p.consecuenciasDeConocerEscenario(e)
	e.aumentarVisitas()
}
