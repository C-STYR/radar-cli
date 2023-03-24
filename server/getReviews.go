package main

import "fmt"

func getReviews(reviewsChan *UrlList, resultsChan *ResultList) {

	FindIndices(reviewsChan)
	for {
		select {
		case url := <-reviewsChan.Urls:
			resultsChan.Enqueue(ParseReview(url))
		default:
			fmt.Println("all done!")
			return
		}
	}
}
