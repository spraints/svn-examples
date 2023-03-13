package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	originURL, err := getOriginURL()
	if err != nil {
		log.Fatal(err)
	}

	token, err := getToken()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("todo: push to %s (token len %d)", originURL, len(token))
}

func getOriginURL() (string, error) {
	// todo - 'git config remote.origin.url'
	return "https://github.com/spraints/svn-examples", nil
}

func getToken() (string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return "", fmt.Errorf("GITHUB_TOKEN must be set in the environment")
	}
	return token, nil
}
