package main

import (
  "fmt"
  "sync"
)

type SafeMap struct {
  m map[string]bool
  mux sync.Mutex
}

func (m *SafeMap) check(s string) bool {
  m.mux.Lock()
  defer m.mux.Unlock()
  return m.m[s]
}

func (m *SafeMap) register(s string) {
  m.mux.Lock()
  m.m[s] = true
  m.mux.Unlock()
}

var done SafeMap

type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  // TODO: Fetch URLs in parallel.
  if depth <= 0 || done.check(url) {
    return
  }

  body, urls, err := fetcher.Fetch(url)
  done.register(url)

  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    Crawl(u, depth-1, fetcher)
  }
  return
}

func main() {
  done = SafeMap{m: make(map[string]bool)}
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
