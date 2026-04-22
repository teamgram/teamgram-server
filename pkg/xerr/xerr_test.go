package xerr

/*
func TestNew(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := New(nil); got != nil {
			t.Fatalf("New(nil) = %v, want nil", got)
		}
	})

	t.Run("wraps with default message", func(t *testing.T) {
		base := errors.New("base error")

		got := New(base)

		assertWrappedError(t, got, base, "base error")
	})
}

func TestNewMsg(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := NewMsg(nil, "custom"); got != nil {
			t.Fatalf("NewMsg(nil) = %v, want nil", got)
		}
	})

	t.Run("wraps with custom message", func(t *testing.T) {
		base := errors.New("base error")

		got := NewMsg(base, "custom message")

		assertWrappedError(t, got, base, "custom message")
	})
}

func TestNewWithErrorLog(t *testing.T) {
	base := errors.New("base error")
	writer := newCaptureWriter()
	restoreLogWriter(t, writer)

	got := NewWithErrorLog(base, "save user %d failed: %v", 1, base)

	assertWrappedError(t, got, base, "base error")
	assertLogs(t, writer.errorLogs(), []string{"save user 1 failed: base error"})
	assertLogs(t, writer.infoLogs(), nil)
}

func TestNewMsgWithErrorLog(t *testing.T) {
	base := errors.New("base error")
	writer := newCaptureWriter()
	restoreLogWriter(t, writer)

	got := NewMsgWithErrorLog(base, "user visible", "save user %d failed: %v", 1, base)

	assertWrappedError(t, got, base, "user visible")
	assertLogs(t, writer.errorLogs(), []string{"save user 1 failed: base error"})
	assertLogs(t, writer.infoLogs(), nil)
}

func TestNewWithInfoLog(t *testing.T) {
	base := errors.New("base error")
	writer := newCaptureWriter()
	restoreLogWriter(t, writer)

	got := NewWithInfoLog(base, "skip sync for user %d: %v", 1, base)

	assertWrappedError(t, got, base, "base error")
	assertLogs(t, writer.infoLogs(), []string{"skip sync for user 1: base error"})
	assertLogs(t, writer.errorLogs(), nil)
}

func TestNewMsgWithInfoLog(t *testing.T) {
	base := errors.New("base error")
	writer := newCaptureWriter()
	restoreLogWriter(t, writer)

	got := NewMsgWithInfoLog(base, "user visible", "skip sync for user %d: %v", 1, base)

	assertWrappedError(t, got, base, "user visible")
	assertLogs(t, writer.infoLogs(), []string{"skip sync for user 1: base error"})
	assertLogs(t, writer.errorLogs(), nil)
}

func assertWrappedError(t *testing.T, got error, wantBase error, wantMsg string) {
	t.Helper()

	if got == nil {
		t.Fatal("got nil error")
	}
	if got.Error() != wantMsg {
		t.Fatalf("Error() = %q, want %q", got.Error(), wantMsg)
	}
	if !errors.Is(got, wantBase) {
		t.Fatalf("errors.Is(%v, %v) = false, want true", got, wantBase)
	}
	if errors.Unwrap(got) != wantBase {
		t.Fatalf("Unwrap() = %v, want %v", errors.Unwrap(got), wantBase)
	}
}

func assertLogs(t *testing.T, got []string, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("log count = %d, want %d; logs=%v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("log[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func restoreLogWriter(t *testing.T, writer logx.Writer) {
	t.Helper()

	prev := logx.Reset()
	logx.SetWriter(writer)
	t.Cleanup(func() {
		logx.Reset()
		if prev != nil {
			logx.SetWriter(prev)
		}
	})
}

type captureWriter struct {
	mu     sync.Mutex
	errors []string
	infos  []string
}

func newCaptureWriter() *captureWriter {
	return &captureWriter{}
}

func (w *captureWriter) Alert(v any)                     {}
func (w *captureWriter) Close() error                    { return nil }
func (w *captureWriter) Debug(v any, _ ...logx.LogField) {}
func (w *captureWriter) Severe(v any)                    {}
func (w *captureWriter) Slow(v any, _ ...logx.LogField)  {}
func (w *captureWriter) Stack(v any)                     {}
func (w *captureWriter) Stat(v any, _ ...logx.LogField)  {}

func (w *captureWriter) Error(v any, _ ...logx.LogField) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.errors = append(w.errors, toString(v))
}

func (w *captureWriter) Info(v any, _ ...logx.LogField) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.infos = append(w.infos, toString(v))
}

func (w *captureWriter) errorLogs() []string {
	w.mu.Lock()
	defer w.mu.Unlock()
	return append([]string(nil), w.errors...)
}

func (w *captureWriter) infoLogs() []string {
	w.mu.Lock()
	defer w.mu.Unlock()
	return append([]string(nil), w.infos...)
}

func toString(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	if err, ok := v.(error); ok {
		return err.Error()
	}
	return ""
}

*/
