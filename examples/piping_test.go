package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestProcessPipe_PipedInput(t *testing.T) {
	// Simulate piped input by replacing os.Stdin with a pipe
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	input := "hello\nworld"
	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	output := captureOutput(ProcessPipe)
	os.Stdin = oldStdin

	assert.Equal(t, "hello\nworld", output)
}
