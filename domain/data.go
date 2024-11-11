package domain

// Entidad que representa los datos scrapeados
type Data struct {
	Title string
	Link  string
}

// Definimos la interfaz del repositorio (base de datos)
type DataRepository interface {
	Save(data []Data) error
	IsDuplicate(link string) (bool, error)
}

// Definimos la interfaz del servicio de scraping
type Scraper interface {
	Scrape(url string) ([]Data, error)
}
