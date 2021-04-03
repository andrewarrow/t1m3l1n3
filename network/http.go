package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"t1m3l1n3/cli"
	"time"
)

func BaseUrl() string {
	url := os.Getenv("CLT_HOST")
	if url == "" {
		return "http://127.0.0.1:8080/"
	}
	return url
}

func DoGet(route string) string {
	agent := "agent"

	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("GET", urlString, nil)
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Username", cli.Username)
	request.Header.Set("Tlz-Server", cli.Server)
	request.Header.Set("Tlz-Index", cli.IndexString)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 500}
	return DoHttpRead("GET", route, client, request)
}

func DoHttpRead(verb, route string, client *http.Client, request *http.Request) string {
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			os.Exit(1)
			return ""
		}
		if resp.StatusCode == 200 || resp.StatusCode == 201 {
			return string(body)
		} else {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, string(body))
			os.Exit(1)
			return ""
		}
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	os.Exit(1)
	return ""
}

func DoPost(uid, username, route string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl(), route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("TLZ-Server", cli.Server)
	request.Header.Set("TLZ-Index", cli.IndexString)
	request.Header.Set("Username", username)
	request.Header.Set("Universe", uid)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead("POST", route, client, request)
}
