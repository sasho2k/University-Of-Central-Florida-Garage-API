package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func Check(err error){
	if err != nil {
		panic(err)
		return
	}
}

func ParseGarageRequest(GarageIdentity string) GarageEntity{
	/*  We take two regex expressions because the site has a return key in there that I didnt want to keep googling.
		Then we get our information, get some more info we might need, and make our object. Easy. In and out.
	*/
	regex := regexp.MustCompile("Garage ([A-Za-z]*)")
	regex2 := regexp.MustCompile("\\d*\\/\\d*")
	regex3 := regexp.MustCompile("percent: (-*[0-9]+)")

	garageName := regex.FindString(GarageIdentity)
	// Garage Numbers is an array of [spaces used / spaces available]
	garageNumbers := strings.Split(regex2.FindString(GarageIdentity),"/")
	garagePercent := strings.Split(regex3.FindString(GarageIdentity), ": ")

	// Quick maths
	currentSpots, _ := strconv.Atoi(garageNumbers[0])
	totalSpots, _ := strconv.Atoi(garageNumbers[1])
	percent, _ := strconv.Atoi(garagePercent[1])

	if percent <= -1 {
		percent = 0
	}

	garage := GarageEntity {
		garageName,
		currentSpots,
		totalSpots-currentSpots,
		totalSpots,
		percent,
	}

	return garage
}

func (g *GarageEntity) print() {
	fmt.Println(g.Name, ":", g.current, "/", g.total, "\t\\\\\t", g.openSpots, "open spots, garage is", g.percent, "% open")
}

func GarageRequest(){
	///////
	request, err := http.Get("https://secure.parking.ucf.edu/GarageCount/")
	Check(err)
	////////

	var garages[7] GarageEntity
	doc, _ := goquery.NewDocumentFromReader(request.Body)
	doc.Find("tr.dxgvDataRow_DevEx").Each(func(i int, s *goquery.Selection) {
		GarageIdentity := s.Find("td").Text()
		garages[i] = ParseGarageRequest(GarageIdentity)
		garages[i].print()
	})

}
func main() {
	/////////////////////////////////////////////
	fmt.Println("UCF GARAGE SCRAPER V1 - SOLAR & SASHO2K")
	/////////////////////////////////////////////

	GarageRequest()
}
