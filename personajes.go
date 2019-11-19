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

type iPersonaje interface {
	felicidad() int
	interactuarConMuchos()
	interactuar()
}

func felicidad(p Personaje) int {
	return felicidad(p)
}

func (h Huesped) felicidad() int {
	return h.minutosRestantes * felicidadDeAmigos(h)
}

func felicidadDeAmigos(h Huesped) int {
	sum := 0
	amigues := h.amigos
	for _, num := range amigues {
		sum += felicidad(num)
	}
	return sum
}

func (a Anfitrion) felicidad() int { // a es como el self
	parque1 := Parque{factor: 10}
	return a.energia / a.velocidadProcesamiento * parque1.factor
}

func (p *Personaje) interactuar() { // (p *Personaje)
	p.energia = p.energia / 2
}

func (a *Anfitrion) interactuar() {
	a.energia = a.energia / 2
	a.velocidadProcesamiento = a.velocidadProcesamiento / 2
}

func (h *Huesped) interactuar() {
	h.energia = h.energia / 2
	h.minutosRestantes = h.minutosRestantes - 10
}

func (p Personaje) rebeldia() int {
	return 1 / felicidad(p)
}

func interactuarConMuchos(personajes []Personaje) {
	/*for _, num := range personajes {
		interactuar(num)
	}*/
}

func consecuenciasDeConocerEscenario(a Anfitrion, e Escenario) {
	a.recuerdos = append(a.recuerdos, Recuerdo{descripcion: "Conoci un escenario", escenario: e})
}

func (h *Huesped) consecuenciasDeConocerEscenario() {
	h.minutosRestantes = h.minutosRestantes - 10
}

func conocerEscenario(p Personaje, e Escenario) {
	p.energia = p.energia - fama(e) - consecuenciasDeConocerEscenario(p, e)
}
