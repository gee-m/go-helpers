package gstrings

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceAnyBetween(t *testing.T) {
	in := "{{++++.Value++}}+++++"
	out := ReplaceAnyBetween(in, "{{", "}}", "+", " ")
	assert.Equal(t, "{{    .Value  }}+++++", out)

	in = "++{{.Value++}}{{++}}+"
	out = ReplaceAnyBetween(in, "{{", "}}", "+", " ")
	assert.Equal(t, "++{{.Value  }}{{  }}+", out)
}

func TestReplaceFuncBetween(t *testing.T) {
	rpl := func(s string) string {
		s, err := url.QueryUnescape(s)
		if err != nil {
			panic(err)
		}
		return s
	}
	in := "asdfas++{{+++%22hello-there%22+++dfa{{+FAds+}}FAS+F}}"
	out := ReplaceFuncBetween(in, "{{", "}}", rpl)
	assert.Equal(t, `asdfas++{{   "hello-there"   dfa{{ FAds }}FAS F}}`, out)
}
