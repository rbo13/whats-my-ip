package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	url := "https://api.ipify.org/?format=json"

	res, err := http.Get(url)

	if err != nil {
		print(err)
	}

	defer res.Body.Close()

	var data map[string]string

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		print(err)
	}

	print(data["ip"])
}
