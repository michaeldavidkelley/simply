package simply_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/michaeldavidkelley/simply"
)

func TestNewLoggerOutput(t *testing.T) {
	old := os.Stderr // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stderr = w

	simply.NewLogger().Info("testing")

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stderr = old // restoring the real stdout
	out := <-outC

	if !strings.Contains(out, `"level":"info"`) {
		t.Fatal("Logger.Info failed to print level")
	}
	if !strings.Contains(out, `"msg":"testing"`) {
		t.Fatal("Logger.Info failed to print msg")
	}
}

func TestLoggerWith(t *testing.T) {
	t.Parallel()

	output := outputWrapper(func(out io.Writer) {
		simply.NewCustomLogger(out).
			With(simply.F{
				"key": 123,
			}).Info("testing")
	})

	simply.AssertContains(t, output, `"key":123`)
}

func TestLoggerError(t *testing.T) {
	t.Parallel()

	output := outputWrapper(func(out io.Writer) {
		simply.NewCustomLogger(out).Error("testing")
	})

	simply.AssertContains(t, output, `"level":"error"`)
}

func TestLoggerErr(t *testing.T) {
	t.Parallel()

	output := outputWrapper(func(out io.Writer) {
		err := errors.New("test")
		simply.NewCustomLogger(out).Err(err).Error("testing")
	})

	simply.AssertContains(t, output, `"err":"test"`)
}

func outputWrapper(fn func(out io.Writer)) string {
	r, w, _ := os.Pipe()

	fn(w)

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()

	return <-outC
}
