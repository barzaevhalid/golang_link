package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest_api/internal/auth"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@mail.ru",
		Password: "test",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("Expected %d got %d", 200, res.StatusCode)
	}

}
