package main

func main() {
  urlChan := CreateURLChannel()
  resultsChan := CreateResultChannel()
  GetReviews(&urlChan, &resultsChan)
}
