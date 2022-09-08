package main

import (
	"testing"
)

func TestGetRandomQuote(t *testing.T) {

	sut := GetRandom()

	if len(sut.Quote) <= 0 {
		t.Errorf("No response body.")
	}

}
