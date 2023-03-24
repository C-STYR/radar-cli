package main

import (
	"fmt"
	"net/http"
)

/*
FindIndices starts at "https://www.undertheradarmag.com/reviews/category/music"
then proceeds to add "P10", "P20", etc in a loop, visiting each site and returning the html
*/
func FindIndices(urlChan *UrlList) {

	urls := []string{
		// "https://www.undertheradarmag.com/reviews/category/music",
		// "https://www.undertheradarmag.com/reviews/category/music/P10",
		"https://www.undertheradarmag.com/reviews/category/music/P20",
		"https://www.undertheradarmag.com/reviews/category/music/P30",
	}

	for _, page := range urls {

		res, err := http.Get(page)
		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		// parse index page and find links to reviews
		links := ReadIndex(res.Body)

		// loop over links map send 
		for url := range links {
			// convert string to ReviewURL and send down channel
			// is type conversion necessary? 
			urlChan.Enqueue(CreateReviewUrl(url))
		}
	}
}
