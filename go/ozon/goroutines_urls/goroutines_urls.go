package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
	}
	wg := sync.WaitGroup{}
	client := &http.Client{}

	outCh := make(chan string)
	errCh := make(chan error)

	const wgInc = 1
	ctx, cancel := context.WithCancel(context.Background())

	var count int
	mtx := sync.Mutex{}
	const outThreshold = 2

	for _, url := range urls {
		url := url
		wg.Add(wgInc)

		go func() {
			defer wg.Done()
			out, err := handleRequest(ctx, client, url)

			outCh <- fmt.Sprint(url, " - ", out)
			errCh <- err

			if out {
				mtx.Lock()
				defer mtx.Unlock()
				count++

				if count == outThreshold {
					cancel()
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
		close(errCh)
	}()

	go func() {
		for err := range errCh {
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	for out := range outCh {
		fmt.Println(out)
	}
}

func handleRequest(ctx context.Context, client *http.Client, requestURL string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == http.StatusOK, nil
}
