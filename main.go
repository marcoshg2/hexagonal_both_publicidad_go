package main

import (
	"encoding/json"
	"hexagonal_both_publicidad_go/adapters"
	"hexagonal_both_publicidad_go/application"
	"log"
	"net/http"
)

func main() {
	// Configurar dependencias
	repo, err := adapters.NewMongoRepository("mongodb://localhost:27017", "scraping_db", "data_collection")
	if err != nil {
		log.Fatal(err)
	}

	scraper := adapters.NewWebScraper()
	service := application.NewScrapeService(scraper, repo)

	http.HandleFunc("/scrape", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "URL requerida", http.StatusBadRequest)
			return
		}

		if err := service.ScrapeAndSave(url); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"message": "Scraping completado"})
	})

	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
