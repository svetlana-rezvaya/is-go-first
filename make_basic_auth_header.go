package main

import (
	"encoding/base64"
	"os"
)

func makeBasicAuthHeader(usernameEnv string, passwordEnv string) string {
	username := os.Getenv(usernameEnv)
	password := os.Getenv(passwordEnv)
	if username == "" || password == "" {
		return ""
	}

	credentials := username + ":" + password
	credentials = base64.StdEncoding.EncodeToString([]byte(credentials))

	return "Basic " + credentials
}
