package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func cacheReviews(rdb *redis.Client) {
	ReviewsToParse := CreateURLChannel()
	Results := CreateResultChannel()

	getReviews(&ReviewsToParse, &Results)

	go func() {
		for {
			select {

			case r := <-Results.Results:
				fmt.Println("got album:", r.albumTitle)
				err := rdb.Set(ctx, r.group, r.albumTitle, 0)
				if err != nil {
					panic(err)
				}
			default:
				fmt.Println("redis write complete")
			}
		}
	}()
}

/*
err := rdb.Set(ctx, "name", "cole", 0).Err()
if err != nil {
	panic(err)
}

val, err := rdb.Get(ctx, "name").Result()
if err != nil {
	panic(err)
}
*/
