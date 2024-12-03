package main

import (
	"fmt"
	"net/http"
)

func urlStatus(url string) {

	for _,url := 
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}

func main() {
	urlStatus("http://google.com")
}
