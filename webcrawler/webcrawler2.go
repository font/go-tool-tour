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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	urlCache.mux.Lock()

	if depth <= 0 || urlCache.isUrlCached(url) {
		urlCache.mux.Unlock()
		wg.Done()
		return
	}

	urlCache.cacheUrl(url)
	urlCache.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}

	wg.Done()
	return
}

func main() {
	wg.Add(1)
	go Crawl("http://golang.org/", 4, fetcher)
	wg.Wait()
}

type fakeUrlCache struct {
	cache map[string]bool
	mux sync.Mutex
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

func (u *fakeUrlCache) cacheUrl(url string) {
	u.cache[url] = true
}

func (u *fakeUrlCache) isUrlCached(url string) bool {
	_, ok := u.cache[url]
	return ok
}

var wg sync.WaitGroup

var urlCache = fakeUrlCache{
	cache: make(map[string]bool),
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

