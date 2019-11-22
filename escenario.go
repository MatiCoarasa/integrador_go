package main

type Escenario struct {
	nombre    string
	categoria ICategoria
}

type ICategoria interface {
	extra() int
	aumentarVisitas()
	evolucion()
}

type BajoCoste struct {
	zona      string
	categoria ICategoria
}

type Estandar struct {
	categoria ICategoria
}

type DeLujo struct {
	visitas   int
	categoria ICategoria
}

func (e Escenario) fama() int {
	return 100 + e.categoria.extra()
}

func (e Escenario) aumentarVisitas() {
	e.categoria.aumentarVisitas()
}

func (e Escenario) evolucionar() {
	e.categoria.evolucion()
}

func (d DeLujo) fama() int {
	return d.visitas
}

func (b BajoCoste) fama() int {
	return len(b.zona)
}

func (e Estandar) fama() int {
	return 10
}

func (d DeLujo) aumentarVisitas() {
	d.visitas = d.visitas + 1
}

func (b BajoCoste) aumentarVisitas() {

}

func (e Estandar) aumentarVisitas() {
}

func (e Escenario) evolucion() {
	e.categoria.evolucionar()
}

func (d DeLujo) extra() int {
	return d.visitas
}

func (b BajoCoste) extra() int {
	return len(b.zona)
}

func (e Estandar) extra() int {
	return 10
}

func (d DeLujo) evolucionar() {

}

func (b BajoCoste) evolucionar() {

}

func (e Estandar) evolucionar() {

}

// Funcion cambiar categoria
