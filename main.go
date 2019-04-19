package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	ip := IP()
	print(ip)
}

// IP returns your current IP address.
func IP() string {
	url := "https://api.ipify.org/?format=json"

	res, err := http.Get(url)

	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()

	var data map[string]string

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err.Error()
	}

	return data["ip"]
}
