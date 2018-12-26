package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		fetch(w)
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func fetch(out io.Writer) {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(out, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Fprintf(out, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
