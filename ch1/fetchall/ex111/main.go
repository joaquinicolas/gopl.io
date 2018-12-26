package main

import (
	"strings"
	"bufio"
	"encoding/csv"
	"sync"
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"time"
	"fmt"
)

const filePath = "./output"
const csvPath = "./top-1m.csv"
var wg sync.WaitGroup

func main() {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: %v\n", csvPath, err)
		os.Exit(1)
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: %v\n", filePath, err)
		os.Exit(1)
	}
	start := time.Now()
	ch := make(chan string)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "reading  %s: %v", csvPath, err)
			os.Exit(1)
		}

		wg.Add(1)
		go fetch(line[1], ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
loop:
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				break loop
			}
			fmt.Fprintln(file, data)
		}
	}
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string)  {
	defer wg.Done()
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("https://%s", url)
	}
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