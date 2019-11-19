package main

type Escenario struct {
	nombre    string
	categoria iCategoria
}

type iCategoria interface {
}

func (e Escenario) fama() int {
	return 100 + extra(e.categoria)
}

func (e Escenario) aumentarVisitas() {
	aumentarVisitas(e.categoria)
}

func (e Escenario) evolucionar() {
	evolucion(e.categoria)
}
