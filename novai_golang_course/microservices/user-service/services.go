package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CallAuthLogin(username, password string) (string, error) {
	requestBody := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
	resp, err := http.Post("http://localhost:8081/login", "application/json", strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
