# RADAR-CLI

## Purpose

This tool is a web-scraper designed to visit the website
[Under the Radar](https://www.undertheradarmag.com/reviews/category/music)
and retrieve information on music reviews found there.

## Process Flow

1. visit music site root
2. parse html looking for links to reviews
3. visit each review link
4. scrape review info
5. A) populate database with individual album reviews - dump all reviews OR
6. B) update database with most recent reviews since last scraped


Currently, this happens through these functions:

main => GetReviews() => FindIndices() => ReadIndex() => ParseReview => print

## Running the tool
1. In its current state, just compile the binary and run on the command line with `./radar-cli`
