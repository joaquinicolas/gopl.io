package main

import (
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"time"
	"fmt"
)

const filePath = "./output"
func main() {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: %v\n", filePath, err)
		os.Exit(1)
	}
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}