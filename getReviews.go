package main

import "fmt"

/* 
GetReviews is the entrypoint - it calls FindIndices() and then listens for review links 
- when links arrive, it calls ParseReview() and sends the review struct down the results channel
*/
func GetReviews(reviewsChan *UrlList, resultsChan *ResultList) {

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
