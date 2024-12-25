package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}
	isDiscard := true
	var count int64
	if isDiscard {
		nbytes, err := io.Copy(io.Discard, resp.Body)
		count = nbytes
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
	} else {
		//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		f, err := os.Create(time.Now().Format(time.RFC3339) + ".test.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		nbytes, err := io.Copy(f, resp.Body)
		count = nbytes
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
	}
	err = resp.Body.Close()
	if err != nil {
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", secs, count, url)
}
