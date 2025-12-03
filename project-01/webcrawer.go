package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

var linkRegex = regexp.MustCompile(`href="(.*?)"`)

func extractLinks(body []byte, base *url.URL) []string {
	matches := linkRegex.FindAllSubmatch(body, -1)
	var links []string
	for _, m := range matches {
		href := string(m[1])
		u, err := url.Parse(href)
		if err == nil {
			resolved := base.ResolveReference(u)
			if resolved.Host == base.Host {
				links = append(links, resolved.String())
			}
		}
	}
	return links
}

func main() {
	startURL := "https://example.com"
	maxDepth := 2

	type task struct {
		url   string
		depth int
	}

	visited := make(map[string]bool)
	var mu sync.Mutex
	tasks := make(chan task, 50)
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for t := range tasks {
			if t.depth > maxDepth {
				continue
			}

			resp, err := http.Get(t.url)
			if err != nil {
				continue
			}
			body := make([]byte, resp.ContentLength)
			resp.Body.Read(body)
			resp.Body.Close()

			base, _ := url.Parse(t.url)
			found := extractLinks(body, base)

			for _, link := range found {
				mu.Lock()
				if !visited[link] {
					visited[link] = true
					mu.Unlock()
					tasks <- task{link, t.depth + 1}
				} else {
					mu.Unlock()
				}
			}
			fmt.Println("Visited:", t.url)
		}
	}

	// start workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker()
	}

	visited[startURL] = true
	tasks <- task{startURL, 0}

	time.AfterFunc(4*time.Second, func() { close(tasks) })
	wg.Wait()
}

