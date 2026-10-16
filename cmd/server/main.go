
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	pages := template.Must(
		template.ParseGlob("templates/pages/*.html"),
	)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),
	))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		err := pages.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			log.Printf("error rendering home: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Sonora running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
