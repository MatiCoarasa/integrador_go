package main

type Escenario struct {
	nombre    string
	categoria ICategoria
}

type ICategoria interface {
	extra() int
	aumentarVisitas()
	evolucionar(es *Escenario)
	fama() int
}

type BajoCoste struct {
	zona string
}

type Estandar struct {
}

type DeLujo struct {
	visitas int
}

func (e Escenario) fama() int {
	return 100 + e.categoria.extra()
}

func (e Escenario) aumentarVisitas() {
	e.categoria.aumentarVisitas()
}

func evolucion(i ICategoria, e *Escenario) {
	i.evolucionar(e)
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

func famaDeEscenario(i ICategoria) int {
	return i.fama()
}

func (d DeLujo) aumentarVisitas() {
	d.visitas = d.visitas + 1
}

func (b BajoCoste) aumentarVisitas() {

}

func (e Estandar) aumentarVisitas() {
}

func (e *Escenario) evolucion(i ICategoria) {
	i.evolucionar(e)
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

func (d DeLujo) evolucionar(es *Escenario) /*error*/ {
	//return errors.New("esta categoria no puede evolucionar")
}

func (b BajoCoste) evolucionar(es *Escenario) {
	(*es).categoria = new(Estandar)
}

func (e Estandar) evolucionar(es *Escenario) {
	(*es).categoria = new(DeLujo)
}
