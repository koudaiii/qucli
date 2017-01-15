package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func HttpPost(url string, apitoken string, body []byte) (string, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if apitoken != "" {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apitoken)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error!\nURL: %s\nstatus code: %d\nbody:\n%s\n", url, resp.StatusCode, string(b))
	}

	return string(b), nil
}
