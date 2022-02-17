// +build integration

package main

import (
	"net/http"
	"testing"
)

func TestIsGoFirst(test *testing.T) {
	type repo struct {
		FullName string `json:"full_name"`
	}
	type repoPage struct {
		Items []repo
	}

	page := repoPage{}
	const url = "https://api.github.com/search/repositories" +
		"?q=programming+language&sort=stars"
	if err := loadJSONData(&http.Client{}, url, &page); err != nil {
		test.Logf("unable to load the data: %s", err)
		test.FailNow()
	}

	if len(page.Items) == 0 {
		test.Log("repos not found")
		test.FailNow()
	}
	if page.Items[0].FullName != "golang/go" {
		test.Log("Go is not first")
		test.FailNow()
	}
}
