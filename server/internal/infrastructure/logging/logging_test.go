package logging_test

import (
	"bytes"
	"context"
	"log"
	"os"
	"room_read/internal/infrastructure/logging"
	"strings"
	"testing"
)

func TestLogFunctions(t *testing.T) {
	ctx := context.Background()

	t.Run("Info", func(t *testing.T) {
		logOutput := captureOutput(func() {
			logging.Info(ctx, "this is an info message")
		})
		assertLogOutputContains(t, logOutput, "INFO this is an info message")
	})
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func assertLogOutputContains(t *testing.T, logOutput, expected string) {
	if !strings.Contains(logOutput, expected) {
		t.Errorf("Expected log output to contain: %s\nActual log output: %s", expected, logOutput)
	}
}
