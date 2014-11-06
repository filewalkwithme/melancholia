package tests

import (
	"github.com/gerep/melancholia"
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
)

func TestPing(t *testing.T) {
	resp := httptest.NewRecorder()
	melancholia.CreateRoutes()

	req, err := http.NewRequest("GET", "http://localhost/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), `I'm ping @ ping`) {
			t.Errorf("header response doen't match:\n%s", p)
		}
	}
}
