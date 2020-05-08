package moviescrawler

import (
	"io/ioutil"
	"net/http"
	"time"
)

func (r WebRequest) Get(Uri string) string {
	req, err := http.NewRequest("GET", Uri, nil)
	if err != nil {
		panic(err)
	}

	for key, value := range r.Header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 120,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
