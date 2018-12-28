package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//	b, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
	//		return
	//	}
	//	err = ioutil.WriteFile(url.QueryEscape(uri), b, 0777)
	//	if err != nil {
	//		ch <- fmt.Sprintf("while writing to %s: %v", uri, err)
	//		return
	//	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%0.2fs  %7d  %s", secs, nbytes, uri)
}
