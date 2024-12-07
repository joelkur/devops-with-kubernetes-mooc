package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port, isPortDefined := os.LookupEnv("PORT")

	if !isPortDefined {
		port = "3000"
	}

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			status := http.StatusMethodNotAllowed
			http.Error(w, http.StatusText(status), status)
		} else {
			tmpl.Execute(w, nil)
		}
	})

	log.Printf("Server started in port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("Failed to start server: %s", err)
		os.Exit(1)
	}
}
