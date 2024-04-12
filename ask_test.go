package main

import (
	"os"
	"strings"
	"testing"
)

func TestReadStdin(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		err      string // Expected error string if any
	}{
		{
			name:     "Simple input",
			input:    "Hello, world!\n",
			expected: "Hello, world!\n",
		},
		{
			name:     "Empty input",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a pipe to simulate stdin
			reader, writer, err := os.Pipe()
			if err != nil {
				t.Fatal("Error creating pipe:", err)
			}

			// Replace os.Stdin with the pipe's reader
			oldStdin := os.Stdin
			os.Stdin = reader

			// Write input to the pipe
			_, err = writer.WriteString(tc.input)
			if err != nil {
				t.Fatal("Error writing to pipe:", err)
			}
			writer.Close() // Close the writer

			// Call the function being tested
			result, err := readStdin()

			// Restore os.Stdin
			os.Stdin = oldStdin

			// Assertions
			if err != nil && !strings.Contains(err.Error(), tc.err) {
				t.Errorf("Unexpected error: got %v, want substring %v", err, tc.err)
			}
			if err == nil && tc.err != "" {
				t.Errorf("Expected an error containing %v, got nil", tc.err)
			}
			if result != tc.expected {
				t.Errorf("Expected output: %q, got: %q", tc.expected, result)
			}
		})
	}
}
