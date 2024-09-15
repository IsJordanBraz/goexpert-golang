package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{Nome: "Go", CargaHoraria: 60}
	template := template.New("cursoTemplate")
	template, _ = template.Parse("Curso: {{.Nome}} - Carga Horaria: {{ .CargaHoraria}}")
	err := template.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
