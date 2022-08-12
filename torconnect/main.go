package main

import (
	"io"
	"os"
	"log"
)

func main() {
	url := ""
	sockClient, err := NewClient("127.0.0.1:9050")

	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		url = os.Args[1] + ":" + os.Args[2]
	}
	
	resp, err := sockClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
