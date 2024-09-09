package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// using defer to close the response body of an HTTP request

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
