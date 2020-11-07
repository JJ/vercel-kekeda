package handler

import (
    "fmt"
    "net/http"
    "time"
)

type Hito struct {
	URI  string
	Title string
	fecha time.Time
}


var hitos = []Hito {
	Hito {
		URI: "0.Repositorio",
		Title: "Datos básicos y repo",
		fecha: time.Date(2020, time.September, 29, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "1.Infraestructura",
		Title: "HUs y entidad principal",
		fecha: time.Date(2020, time.October, 6, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "2.Tests",
		Title: "Tests iniciales",
		fecha: time.Date(2020, time.October, 16, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "3.Contenedores",
		Title: "Contenedores",
		fecha: time.Date(2020, time.October, 26, 11, 30, 0, 0, time.UTC),
	},
	Hito {
		URI: "4.CI",
		Title: "Integración continua",
		fecha: time.Date(2020, time.November, 6, 23, 59, 0, 0, time.UTC),
	},
	Hito {
		URI: "5.Microservicios",
		Title: "Trabajando con microservicios",
		fecha: time.Date(2020, time.November, 17, 11, 30, 0, 0, time.UTC),
	},

}


func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	var next int
	var queda time.Duration
	for indice, hito := range hitos {
		if ( hito.fecha.After( currentTime ) ) {
			next = indice
			queda = hito.fecha.Sub( currentTime )
		}
	}
	if ( next > 0 ) {
		fmt.Fprintf(w, queda.String())
	} else {
		fmt.Fprintf(w, "Ninguna entrega próxima" )
	}
}
