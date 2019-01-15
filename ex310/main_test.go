package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComma2(t *testing.T) {
	testcases := []struct {
		in   string
		want string
	}{
		{"0", "0"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
	}

	for _, tc := range testcases {
		s := comma2(tc.in)
		assert.Equal(t, tc.want, s)
	}
}
