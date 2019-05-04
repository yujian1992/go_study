package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{IndentJSON: true})
)

const (
	fackMatchLocationResult = "/matches/t37eu78-3245-3456-de45-hsjj876et5gdi"
)

func TestTestHandler(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(testHandler(formatter)))
	defer server.Close()

	body := []byte("{\n \"gridsize\":19,\n \"players\": [\n {\n \"color\": \"whiite\",\n \"name\": \"bob\"\n },\n {\n \"color\": \"black\",\n \"name\": \"alff\"\n }\n ]\n}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in create POST:%v", err)
	}
	req.Header.Add("Content-Type", "aplication/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in post:%v", err)
	}

	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error readall:%v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201,recvived:%v", err)
	}
	fmt.Printf("payload: %s", string(payload))

	loc, ok := res.Header["Location"]
	if !ok {
		t.Errorf("Location header is not set")
	} else {
		if !strings.Contains(loc[0], "/matches") {
			t.Errorf("location header muster contian '/matcher',header:%v", loc[0])
		}
		if len(loc[0]) != len(fackMatchLocationResult) {
			t.Errorf("location value dos not contain guid of new match")
		}
	}
}
