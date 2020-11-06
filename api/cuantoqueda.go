package handler

import (
    "fmt"
    "net/http"
    "time"
)

type Hito struct {
	URI  string
	Title string
	Date time.Time
}


var hitos = []Hito {
	Hito {
		URI: "0.Repositorio",
		Title: "Datos básicos y repo",
		Date: time.Date(2020, time.September, 29, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "1.Infraestructura",
		Title: "HUs y entidad principal",
		Date: time.Date(2020, time.October, 6, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "2.Tests",
		Title: "Tests iniciales",
		Date: time.Date(2020, time.October, 16, 11, 30, 0, 0, time.UTC),
	},
}


func Handler(w http.ResponseWriter, r *http.Request) {
	var currentTime Time := time.Now()
	fmt.Println( currentTime )
	var next
	for indice, hito := range hitos {
		if ( hito.Date.after( currentTime ) ) {
			next = indice
		}
	}
	fmt.Fprintf(w, currentTime)
}
