package helloworld

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestF(t *testing.T) {
        data := url.Values{}
        data.Set("message", "Hello World!")
  
	r, err := http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(F)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "Hello World!" {
		t.Errorf("wrong response body: got %v want %v", body, "Hello World!")
	}
}
