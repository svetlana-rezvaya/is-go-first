package main

import (
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

	request = request
}
