package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	visited map[string]bool
	found   []string
	mux     sync.Mutex
}

func (c *SafeCache) Visited(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	if _, ok := c.visited[url]; ok {
		return true
	}
	c.visited[url] = true
	return false

}
func (c *SafeCache) AddFound(url string) {
	c.mux.Lock()
	c.found = append(c.found, url)
	c.mux.Unlock()
}

func (c *SafeCache) GetFound() []string {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.found
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		visited: make(map[string]bool),
		found:   make([]string, 0),
	}
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, c *SafeCache) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url, cache)
	if err != nil {
		fmt.Println(err)
		return
	}
	cache.AddFound(url) // Store the found URL
	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, cache, wg)
	}
}

func mainCrawler() {
	var wg sync.WaitGroup
	cache := NewSafeCache()
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, cache, &wg)
	wg.Wait()
	foundUrls := cache.GetFound()
	fmt.Println("\nAll found URLs:")
	for _, url := range foundUrls {
		fmt.Println(url)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, c *SafeCache) (string, []string, error) {
	if c.Visited(url) {
		return "", nil, fmt.Errorf("not found: %s", url)
	}
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
