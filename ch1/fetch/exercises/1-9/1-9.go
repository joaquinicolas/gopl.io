package main

import (
	"strings"
	"fmt"
	"os"
	"io"
	"log"
	"net/http"
)

func main()  {
	handler := func (w http.ResponseWriter, r *http.Request)  {
		fetch(w)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func fetch(out io.Writer)  {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = fmt.Sprintf("http://%s", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("Status: %s\n", resp.Status)
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url ,err)
			os.Exit(1)
		}
	}
}