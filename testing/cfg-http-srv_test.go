package main

import (
	"net/http"
	"testing"
)

type FakeConfig struct{}
type FakeResponseWriter struct {
	h    http.Header
	Body []byte
}

const (
	msg = "Please send help, I'm trapped in the web server"
)

func (c FakeConfig) Get(key string) (string, error) {
	return msg, nil
}

// It always works!  Nice
func (c FakeConfig) Set(key, val string) error {
	return nil
}

func (wr FakeResponseWriter) Header() http.Header {
	return wr.h
}

func (wr FakeResponseWriter) Write(b []byte) (int, error) {
	wr.Body = b
	return len(msg), nil
}

func (wr FakeResponseWriter) WriteHeader(i int) {}

func TestResponderHandler(t *testing.T) {
	responder := Responder{
		Cfg: FakeConfig{},
	}
	w := FakeResponseWriter{
		h: http.Header{},
	}

	// Don't even care about the request!
	// But if needed to we could control that pretty well too.
	responder.Handler(w, nil)
	header := w.Header().Get("X-Config-Option")
	t.Log(header)
	if header != msg {
		t.Fatalf("Expected X-Config-Option to be %q, got %q", msg, header)
	}
}
