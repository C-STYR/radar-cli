package main

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type Ratings struct {
	group        string
	albumTitle   string
	authorRating float64
	readerRating float64
	// category string
	// reviewLink   ReviewUrl
}

// parseReview parses an html file and returns a Ratings struct of collected info
func ParseReview(url ReviewUrl) Ratings {
	var ratings Ratings

	// ratings.reviewLink = url
	res, err := http.Get(string(url))
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	fileScanner := bufio.NewScanner(res.Body)
	fileScanner.Split(bufio.ScanLines)

	findGroup := regexp.MustCompile(`<h3>([0-9\w\s\.\\!\*\-_:;'",’]+)</h3>`)
	findAlbumTitle := regexp.MustCompile(`<h4><i>([0-9\w\s\.\\!\*\-_:;'",’]+)</i></h4>`)
	findAuthorRating := regexp.MustCompile(`Author rating: <b>([0-9\.]+)</b>`)
	findReaderRating := regexp.MustCompile(`reader rating: <b>([0-9\.]+)</b>`)

	for fileScanner.Scan() {
		Group := findGroup.FindStringSubmatch(fileScanner.Text())
		aTitle := findAlbumTitle.FindStringSubmatch(fileScanner.Text())
		aRating := findAuthorRating.FindStringSubmatch(fileScanner.Text())
		rRating := findReaderRating.FindStringSubmatch(fileScanner.Text())

		// this regexp has several possible matches per page - target the first one only
		if ratings.group == "" && len(Group) > 1 {
			group := Group[1]
			ratings.group = group
		}

		if len(aTitle) > 1 {
			title := aTitle[1]
			ratings.albumTitle = title
		}

		if len(aRating) > 1 {
			num, err := strconv.ParseFloat(aRating[1], 32)
			if err != nil {
				fmt.Println(err)
			}
			ratings.authorRating = num
		}

		if len(rRating) > 1 {
			num, err := strconv.ParseFloat(rRating[1], 32)
			if err != nil {
				fmt.Println(err)
			}
			ratings.readerRating = num
		}
	}
	// printing is not intended behavior - just for testing purposes
	// can tie this into dependency injection model? 
	fmt.Println(ratings)
	return ratings
}
