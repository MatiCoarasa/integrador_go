package main

type Recuerdo struct {
	descripcion string
	escenario   Escenario
}

func (r Recuerdo) emotividad() int {
	return len(r.descripcion) //* fama(r.escenario)
}
