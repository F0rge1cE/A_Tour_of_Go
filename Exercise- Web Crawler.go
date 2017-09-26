package main

import (
	"fmt"
	"sync"
	"errors"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type fetched struct{
	fetched_list map[string]error
	mux sync.Mutex
}

var loading = errors.New("url load in progress") // sentinel value
var fl = fetched{fetched_list: make(map[string]error)};

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	
	if depth <= 0 {
		return
	}	
	
	fl.mux.Lock()
	// Whether this page has already been downloaded. 
	if _,exist_flag:=fl.fetched_list[url]; exist_flag {
		fl.mux.Unlock()
		return
	}
	
	fl.fetched_list[url] = loading
	fl.mux.Unlock()
	
	body, urls, err := fetcher.Fetch(url)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Not necessary in this case.
	/*
	fl.mux.Lock()
	fl.fetched_list[url] = err
	fl.mux.Unlock()
	*/
	
	fmt.Printf("found: %s %q\n", url, body)
	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
		
		<-done
	}
	
	// for range urls {
	// 	 <-done
	// }
	
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
