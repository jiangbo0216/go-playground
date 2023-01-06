package http_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/net/http/request_test.go;bpv=1;bpt=0
func TestQuery(t *testing.T) {
	req := &http.Request{Method: "GET"}

	req.URL, _ = url.Parse("http://baidu.com/search?q=foo&q=bar")

	if q := req.FormValue("q"); q != "foo" {
		t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
	}
}

func TestParseFormQuery(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://baidu.com/search?q=foo&q=bar&both=x&prio=1&orphan=nope&empty=not",
		strings.NewReader("z=post&both=y&prio=2&=nokey&orphan&empty=&"))

	if q := req.FormValue("q"); q != "foo" {
		t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
	}
}
