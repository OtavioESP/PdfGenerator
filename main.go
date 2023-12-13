package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// As chaves para serem identificadas no HTML precisa começar com letra maiuscula
type Historico struct {
	Disciplina   string
	CargaHoraria int
}

func index(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	historicoEscolar := map[string][]Historico{
		"Historico": {
			{Disciplina: "Anatomia Humana", CargaHoraria: 120},
			{Disciplina: "Anatomia Humana", CargaHoraria: 120},
			{Disciplina: "Anatomia Humana", CargaHoraria: 120},
		},
	}

	templ.Execute(w, historicoEscolar)
}

func main() {
	fmt.Println("Hello World!")

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
