package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myh myHandler

	h := NoSurf(&myh)
	switch h.(type) {
	case http.Handler:
	default:
		t.Error("type is not http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var myh myHandler

	h := SessionLoad(&myh)
	switch h.(type) {
	case http.Handler:
	default:
		t.Error("type is not http.Handler")
	}
}
