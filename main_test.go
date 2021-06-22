package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestIsGoFirst(test *testing.T) {
	const url = "https://api.github.com/search/repositories" +
		"?q=programming+language&sort=stars"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		test.Logf("unable to create the request: %s", err)
		test.FailNow()
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		test.Logf("unable to send the request: %s", err)
		test.FailNow()
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		test.Logf("unable to read the response: %s", err)
		test.FailNow()
	}

	type repo struct {
		FullName string `json:"full_name"`
	}
	type repoPage struct {
		Items []repo
	}

	responseData := repoPage{}
	err = json.Unmarshal(responseBytes, &responseData)
	if err != nil {
		test.Logf("unable to unmarshal the response: %s", err)
		test.FailNow()
	}

	if len(responseData.Items) == 0 {
		test.Log("repos not found")
		test.FailNow()
	}
	if responseData.Items[0].FullName != "golang/go" {
		test.Log("Go is not first")
		test.FailNow()
	}
}
