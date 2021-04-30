package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Check :: Checks for error.
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
	currentSpots, err := strconv.Atoi(garageNumbers[0])
	Check(err) // Check if there is an error with the first value, if not we ensure that the array does not contain any nil values.
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

// print :: Debug to print values.
func (g *GarageEntity) print() {
	fmt.Println(g.Name, ":", g.Current, "/", g.Total, "\t\\\\\t", g.OpenSpots, "open spots, garage is", g.Percent, "% open")
}

// GarageRequest :: Creates a request to the resource and parses the information into struct values (GarageEntity).
func GarageRequest() []GarageEntity{

	var requestURL string
	var garages []GarageEntity
	garages = make([]GarageEntity, 7)

	/* URL used for all requests. */
	requestURL = "https://secure.parking.ucf.edu/GarageCount/"
	request, err := http.Get(requestURL)
	Check(err)

	/* Load the request body into goQuery and find/parse the garageBlob into it's struct value (garages[i]). */
	doc, _ := goquery.NewDocumentFromReader(request.Body)
	doc.Find("tr.dxgvDataRow_DevEx").Each(func(i int, s *goquery.Selection) {
		garageBlob := s.Find("td").Text()
		garages[i] = ParseGarageRequest(garageBlob)
	})

	return garages
}
