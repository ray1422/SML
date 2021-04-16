package utils

import "testing"

func Test_QuoteRM(t *testing.T) {
	if QuoteRM(`'www'`) != "www" {
		t.Error("should be `www`, got", QuoteRM(`'www'`))
	}
	if QuoteRM(`"www"`) != "www" {
		t.Error("should be `www`, got", QuoteRM(`"www"`))
	}
	if QuoteRM(`www`) != "www" {
		t.Error("should be `www`, got", QuoteRM(`www`))
	}
	if QuoteRM(`a`) != "a" {
		t.Error("should be `a`, got", QuoteRM(`a`))
	}
}
