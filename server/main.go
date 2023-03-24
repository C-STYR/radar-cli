package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
  - start at https://www.undertheradarmag.com/reviews/category/music
  - next page is https://www.undertheradarmag.com/reviews/category/music/P10
  - visit each page in a goroutine and search the source code for regex matches to
    "https://www.undertheradarmag.com/reviews/*" where * excludes "category"
  - each will be inside a div cn="headline", inside an h3
  - there will be repeats
  - build a map of those links, which will then be traversed to look for ratings
  - then filter for author ratings or reader ratings > 8.5
  - for each review page with a rating greater than 8.5, make an object with the two ratings, the name of the album, the name of the band, a link to the album review page
  - create a slice of objects, serialize and send to the front end
*/
var ctx = context.Background()

func main() {
	// start redis in a terminal with "redis-server"
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cacheReviews(rdb)

	time.Sleep(40 * time.Second)
	fmt.Println("did it work?")
}
