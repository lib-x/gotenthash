package gotenthash

import (
	"bytes"
	"strings"
	"testing"
)

func TestTentHash(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{"Hello world!", []byte{0x15, 0x5f, 0x0a, 0x35}},
		{"I love golang!", []byte{0xf9, 0x8c, 0x95, 0xae}},
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

func TestTentHasherWriteReader(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{"Hello world!", []byte{0x15, 0x5f, 0x0a, 0x35}},
		{"I love golang!", []byte{0xf9, 0x8c, 0x95, 0xae}},
	}

	for _, tc := range testCases {
		hasher := New()
		reader := strings.NewReader(tc.input)

		n, err := hasher.WriteReader(reader)
		if err != nil {
			t.Errorf("WriteReader error: %v", err)
		}
		if n != int64(len(tc.input)) {
			t.Errorf("WriteReader wrote %d bytes, expected %d", n, len(tc.input))
		}

		result := hasher.Sum(nil)

		if !bytes.Equal(result[:4], tc.expected) {
			t.Errorf("For input '%s': Expected %x, got %x", tc.input, tc.expected, result[:4])
		}
	}
}

func TestSumReader(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{"Hello world!", []byte{0x15, 0x5f, 0x0a, 0x35}},
		{"I love golang!", []byte{0xf9, 0x8c, 0x95, 0xae}},
	}

	for _, tc := range testCases {
		reader := strings.NewReader(tc.input)
		h := New()
		result, err := h.SumReader(reader)
		if err != nil {
			t.Errorf("SumReader error: %v", err)
		}

		if !bytes.Equal(result[:4], tc.expected) {
			t.Errorf("For input '%s': Expected %x, got %x", tc.input, tc.expected, result[:4])
		}

		// Compare with Hash function result
		hashResult := Hash([]byte(tc.input))
		if !bytes.Equal(result, hashResult[:]) {
			t.Errorf("SumReader and Hash results differ for input '%s'", tc.input)
		}
	}
}
