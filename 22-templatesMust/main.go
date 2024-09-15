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
	t := template.Must(template.New("cursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria: {{ .CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
