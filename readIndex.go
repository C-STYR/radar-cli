package main

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// TODO: why is links a map and not []string? is it to avoid re-allocation cost of appending to slice? 

// ReadIndex parses an html file and returns a map of urls that match a given regexp
func ReadIndex(r io.Reader) map[string]bool {
	links := make(map[string]bool)

	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanLines)
	reviewLink := regexp.MustCompile(`https://[a-z\.]*/reviews/[A-Za-z0-9\-\._~:\/?#\[\]@!$'&\(\)\*+,;=]*`)

	for fileScanner.Scan() {
		link := reviewLink.FindString(fileScanner.Text())

		/*
			we need to avoid the pattern "...undertheradarmag.com/reviews/category"
			but grab the pattern "...undertheradarmap.com/reviews/(anything else here)"
			golang regexp has no negative lookahead support, so we are manually checking for "category"
			TODO: refine regexp or look for another solution
		*/
		if strings.Contains(link, "category") {
			link = ""
		}
		if len(link) != 0 {
			links[link] = true
		}
	}
	return links
}
