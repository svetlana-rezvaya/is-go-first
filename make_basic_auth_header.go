package main

import "encoding/base64"

func makeBasicAuthHeader() string {
	username := ""
	password := ""

	credentials := username + ":" + password
	credentials = base64.StdEncoding.EncodeToString([]byte(credentials))

	return "Basic " + credentials
}
