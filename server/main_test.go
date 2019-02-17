package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be OK, but received %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	// read a string like below
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// read and unmarsall json
	var user User
	err = json.Unmarshal(b, &user)

	if err != nil {
		t.Fatal(err)
	}

	// if the output is string
	// respString := string(b)

	expected := "Some_User_Id"

	if user.UserId != expected {
		t.Errorf("Response should be %s, but received %s", expected, user.UserId)
	}
}
