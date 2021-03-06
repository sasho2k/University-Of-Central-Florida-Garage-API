package job

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sasho2k/University-Of-Central-Florida-Garage-API/internals"
	"net/http"
)

// GarageRequest :: Creates a request to the resource and parses the information into struct values (GarageEntity).
func GarageRequest() (garages []internals.GarageEntity, err error) {
	garages = make([]internals.GarageEntity, 7)

	/* URL used for all requests. */
	requestURL := "https://secure.parking.ucf.edu/GarageCount/"
	request, err := http.Get(requestURL)

	if err != nil {
		return garages, err
	}

	/* Load the request body into goQuery and find/parse the garageBlob into it's struct value (garages[i]). */
	doc, _ := goquery.NewDocumentFromReader(request.Body)
	doc.Find("tr.dxgvDataRow_DevEx").Each(func(i int, s *goquery.Selection) {
		garageBlob := s.Find("td").Text()
		garages[i], err = internals.ParseGarageRequest(garageBlob)
	})

	return garages, err
}
