package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseGarageRequest(GarageIdentity string) (garage GarageEntity, err error) {
	/*  We take two regex expressions because the site has a return key in there that I didnt want to keep googling.
	Then we get our information, get some more info we might need, and make our object. Easy. In and out.
	*/

	regex := regexp.MustCompile("Garage ([A-Za-z]*)")
	regex2 := regexp.MustCompile("\\d*\\/\\d*")
	regex3 := regexp.MustCompile("percent: (-*[0-9]+)")

	garageName := regex.FindString(GarageIdentity)
	// Garage Numbers is an array of [spaces used / spaces available]
	garageNumbers := strings.Split(regex2.FindString(GarageIdentity), "/")
	garagePercent := strings.Split(regex3.FindString(GarageIdentity), ": ")

	// Quick maths
	currentSpots, err := strconv.Atoi(garageNumbers[0])
	if err != nil {
		return GarageEntity{}, err
	} // Check if there is an error with the first value, if not we ensure that the array does not contain any nil values.
	totalSpots, _ := strconv.Atoi(garageNumbers[1])
	percent, _ := strconv.Atoi(garagePercent[1])

	if percent <= -1 {
		percent = 0
	}

	garage = GarageEntity{
		garageName,
		currentSpots,
		totalSpots - currentSpots,
		totalSpots,
		percent,
	}

	return garage, err
}

func ParseDate(year int, month time.Month, day int, hour int, min int, sec int) string {
	return strconv.Itoa(year) + "/" + month.String() + "/" + strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":" +
		strconv.Itoa(min) + ":" + strconv.Itoa(sec) + "\n"
}

// Print :: Debug to print values.
func (g *GarageEntity) Debug() {
	fmt.Println(g.Name, ":", g.Current, "/", g.Total, " \\\\\t", g.OpenSpots, "open spots, garage is", g.Percent, "% open")
}

// Print :: Debug to print values.
func (g *GarageEntity) Print() string {
	return g.Name + ": " + strconv.Itoa(g.Current) + "/" + strconv.Itoa(g.Total) + "\n        " + strconv.Itoa(g.OpenSpots) +
		" open spots, garage is " + strconv.Itoa(g.Percent) + "% open"
}
