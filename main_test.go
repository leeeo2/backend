package main

import (
	"encoding/json"
	"fmt"
	"github.com/leeeo2/backend/pkg/domain"
	"io"
	"net/http"
	"strings"
	"testing"
)

func request(t *testing.T, url string, input interface{}, header ...interface{}) {
	// marshal input
	js, err := json.Marshal(input)
	if err != nil {
		t.Log("marshal input failed,err:", err)
		return
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(js)))
	if err != nil {
		t.Log("new request failed,err", err)
		return
	}

	// set header
	req.Header.Set("Content-type", "application/json")
	if len(header)%2 != 0 {
		t.Log("header is not valid")
	} else {
		for i := 0; i < len(header); i += 2 {
			key := header[i]
			value := header[i+1]
			req.Header.Set(key.(string), value.(string))
		}
	}

	// do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Log("do request failed,err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp failed,err:", err)
		return
	}

	fmt.Println("resp:", string(b))

}

func TestRegister(t *testing.T) {
	target := "http://localhost:8888/api/auth/register"

	input := domain.RegisterInput{
		Name:      "lxx",
		Telephone: "12345678914",
		Password:  "13245678",
	}
	request(t, target, input)
}

func TestLogin(t *testing.T) {
	target := "http://localhost:8888/api/auth/login"

	input := domain.LoginInput{
		Telephone: "12345678914",
		Password:  "13245678",
	}
	request(t, target, input)
}

func TestUserInfo(t *testing.T) {
	target := "http://localhost:8888/api/auth/info"

	key := "Authorization"
	value := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjQsImV4cCI6MTY5OTI2OTY5MiwiaWF0IjoxNjk4NjY0ODkyLCJpc3MiOiJseCIsInN1YiI6InVzZXIgdG9rZW4ifQ.ae-LNmun0loOD-eZFBVHtRtluJk1l5DEtiuSGPkQvX0"
	request(t, target, nil, key, value)
}
