package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fetchURL(ctx context.Context, url string) {
	defer wg.Done()
	ch := make(chan string)
	go func() {
		http.Get(url)
		ch <- "Fetching done"
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout , url ureachable")
		return

	case <-ch:
		log.Println("finished fetching ", url)
	}

}

func main() {

	urls := []string{"google.com", "youtube.com", "http://httpbin.org/delay/5"}

	for _, url := range urls {
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		wg.Add(1)
		go fetchURL(ctxWithTimeout, url)

	}
	wg.Wait()
}
