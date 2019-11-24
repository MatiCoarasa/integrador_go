package main

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

/*func felicidadDeUnPersonaje(i IPersonaje) int { // implementacion de interfaz
	return i.felicidad()
}*/

func felicidadPersonaje(i IPersonaje) int {
	return i.felicidad()
}

func (h Huesped) felicidad() int {
	return h.minutosRestantes * h.felicidadDeAmigos()
}

func (h Huesped) felicidadDeAmigos() int {
	felicidadTotalDeAmigos := 0
	amigosDelPersonaje := h.amigos
	for _, IPersonaje := range amigosDelPersonaje {
		felicidadTotalDeAmigos += felicidadPersonaje(&IPersonaje)
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

func interaccion(i IPersonaje) { // implementacion de interfaz
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

func (p Personaje) rebeldia() int {
	return 1 / felicidadDeUnPersonaje(&p)
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

func (e Escenario) consecuencias(i IPersonaje) { // implementacion de interfaz
	i.consecuenciasDeConocerEscenario(e)
}

func (p *Personaje) consecuenciasDeConocerEscenario(e Escenario) {}

func (p *Personaje) conocerEscenario(e Escenario) {
	(*p).energia = p.energia - e.fama()
	p.consecuenciasDeConocerEscenario(e)
	e.aumentarVisitas()
}
