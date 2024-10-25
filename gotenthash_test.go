package gotenthash

import (
	"bytes"
	"testing"
)

func TestTentHash(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{"Hello world!", []byte{0x15, 0x5f, 0x0a, 0x35}},
	}

	for _, tc := range testCases {
		result := Hash([]byte(tc.input))
		if !bytes.Equal(result[:4], tc.expected) {
			t.Errorf("For input '%s': Expected %x, got %x", tc.input, tc.expected, result[:4])
		}
	}
}

func TestTentHasher(t *testing.T) {
	expected := []byte{0x15, 0x5f, 0x0a, 0x35}

	hasher := New()
	hasher.Write([]byte("Hello "))
	hasher.Write([]byte("world!"))
	result := hasher.Sum(nil)

	if !bytes.Equal(result[:4], expected) {
		t.Errorf("Expected %x, got %x", expected, result[:4])
	}
}
