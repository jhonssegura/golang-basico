package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) marcarCompleta() {
	t.completado = true
}

func (t task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := task{
		nombre:      "Curso Práctico de Go.",
		descripcion: "Completar en esta Semana.",
	}
	fmt.Printf("%+v\n", t) // Imprime la interfaz y seguir el patron propiedad/valor
	t.marcarCompleta()
	t.actualizarNombre("Finalizar mi curso de Go")
	t.actualizarDescripcion("COmpletar mi curso cuanto antes")
}
