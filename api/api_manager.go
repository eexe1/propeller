package api

import (
	"io/ioutil"
	"net/http"
	"time"
)


func Get(url string) (string, error) {
	tr := &http.Transport{
		MaxIdleConns:		10,
		IdleConnTimeout:	30 * time.Second,
	}

	client := &http.Client{Transport: tr}
	// for testing
	// resp, err := client.Get(Base_api_url + url)
	resp, err := client.Get(url)
	if err != nil {
		return  "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return bodyString, err2
	}

	return "", nil
}
