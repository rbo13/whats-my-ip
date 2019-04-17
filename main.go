package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://httpbin.org/ip"

	res, err := http.Get(url)

	if err != nil {
		print(err)
	}

	defer res.Body.Close()

	byt, err := ioutil.ReadAll(res.Body)

	if err != nil {
		print(err)
	}

	print(string(byt))
}
