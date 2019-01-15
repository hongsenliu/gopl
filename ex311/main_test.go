package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma(t *testing.T) {
	testcases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"0", "0"},
		{"-234.34568", "-234.34568"},
		{"-1", "-1"},
		{"-1345.2345", "-1,345.2345"},
		{"+12345.23", "+12,345.23"},
	}
	for _, tc := range testcases {
		s := comma(tc.in)
		assert.Equal(t, tc.want, s)
	}
}
