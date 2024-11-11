package adapters

import (
	"hexagonal_both_publicidad_go/domain"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type WebScraper struct{}

func NewWebScraper() *WebScraper {
	return &WebScraper{}
}

func (ws *WebScraper) Scrape(url string) ([]domain.Data, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var data []domain.Data
	doc.Find("h2.title").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Find("a").Attr("href")
		data = append(data, domain.Data{Title: title, Link: link})
	})
	return data, nil
}
