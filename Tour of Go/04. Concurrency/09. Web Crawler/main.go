package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type urlData struct {
	body, url string
}

type tracker struct {
	mutex          sync.Mutex
	remainingLinks int
	visitedLinks   map[string]bool
}

func (t *tracker) markVisited() int {
	t.mutex.Lock()
	t.remainingLinks -= 1
	defer t.mutex.Unlock()
	return t.remainingLinks
}

func (t *tracker) addLinks(n int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.remainingLinks += n
}

func (t *tracker) checkVisited(url string) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if _, visited := t.visitedLinks[url]; visited {
		return true
	}

	t.visitedLinks[url] = true
	return false

}

func Crawl(url string, depth int, fetcher Fetcher) {
	ch := make(chan urlData)

	t := &tracker{sync.Mutex{}, 1, make(map[string]bool)}

	go crawlHelper(url, depth, fetcher, ch, t)
	for result := range ch {
		fmt.Printf("body:- %v, url:- %v\n", result.body, result.url)
	}
}

func crawlHelper(url string, depth int, fetcher Fetcher, ch chan urlData, t *tracker) {
	defer func() {
		remaining := t.markVisited()
		if remaining == 0 {
			fmt.Println("Closing Channel, Exiting Program")
			close(ch)
		}
	}()

	if depth <= 0 || t.checkVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	t.addLinks(len(urls))
	ch <- urlData{body, url}
	for _, u := range urls {
		go crawlHelper(u, depth-1, fetcher, ch, t)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
