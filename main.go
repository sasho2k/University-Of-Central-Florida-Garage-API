package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "log"
	"net/http"
)

func Check(err error){
	if err != nil {
		panic(err)
		return
	}
}

func ParseGarageRequest(GarageIdentity string) {

	fmt.Println(GarageIdentity)
}



func GarageRequest(){
	///////
	request, err := http.Get("https://secure.parking.ucf.edu/GarageCount/")
	Check(err)
	////////

	doc, _ := goquery.NewDocumentFromReader(request.Body)
	doc.Find("tr.dxgvDataRow_DevEx").Each(func(i int, s *goquery.Selection) {
		GarageIdentity := s.Find("td").Text()
		if i == 1{
			ParseGarageRequest(GarageIdentity)
		}
	})
}
func main() {
	/////////////////////////////////////////////
	fmt.Println("UCF GARAGE SCRAPER V1")
	/////////////////////////////////////////////

	GarageRequest()
}
