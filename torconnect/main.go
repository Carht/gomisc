package main

import (
	"io"
	"os"
	"log"
)

func main() {
	sockClient, err := NewClient("127.0.0.1:9050")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := sockClient.Get("https://check.torproject.org:443")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, resp.Body)
}
