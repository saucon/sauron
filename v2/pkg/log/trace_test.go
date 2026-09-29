package log

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/saucon/sauron/v2/pkg/log/logconfig"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func newTestLogger(buf *bytes.Buffer) *LogCustom {
	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{})
	lg.SetOutput(buf)
	lg.ExitFunc = func(int) {} // let Fatal/FatalCtx return in tests
	return &LogCustom{Logrus: lg, WhoAmI: &logconfig.Iam{}}
}

func lastEntry(t *testing.T, buf *bytes.Buffer) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		t.Fatalf("decode log line %q: %v", buf.String(), err)
	}
	buf.Reset()
	return m
}

func spanCtx(t *testing.T) (context.Context, trace.SpanContext) {
	t.Helper()
	tid, _ := trace.TraceIDFromHex("0af7651916cd43dd8448eb211c80319c")
	sid, _ := trace.SpanIDFromHex("b7ad6b7169203331")
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: tid, SpanID: sid, TraceFlags: trace.FlagsSampled,
	})
	return trace.ContextWithSpanContext(context.Background(), sc), sc
}

func TestCtxMethods_AddTraceFieldsAndKeepCaller(t *testing.T) {
	gcpProjectID.Store("my-project")
	resolveProjectOnce.Do(func() {}) // skip env/metadata lookup in tests

	var buf bytes.Buffer
	l := newTestLogger(&buf)
	ctx, sc := spanCtx(t)
	data := LogData{Description: "d", Err: errors.New("boom"), Message: "m"}

	l.InfoCtx(ctx, data)
	info := lastEntry(t, &buf)
	l.ErrorCtx(ctx, data)
	errEntry := lastEntry(t, &buf)
	l.AlertCtx(ctx, data)
	alert := lastEntry(t, &buf)
	l.SuccessCtx(ctx, data)
	success := lastEntry(t, &buf)
	l.FatalCtx(ctx, data)
	fatal := lastEntry(t, &buf)

	entries := map[string]map[string]any{
		"InfoCtx": info, "ErrorCtx": errEntry, "AlertCtx": alert, "SuccessCtx": success, "FatalCtx": fatal,
	}
	for name, m := range entries {
		if m["trace_id"] != sc.TraceID().String() || m["span_id"] != sc.SpanID().String() {
			t.Errorf("%s: trace_id/span_id = %v/%v", name, m["trace_id"], m["span_id"])
		}
		if want := "projects/my-project/traces/" + sc.TraceID().String(); m["logging.googleapis.com/trace"] != want {
			t.Errorf("%s: logging.googleapis.com/trace = %v, want %s", name, m["logging.googleapis.com/trace"], want)
		}
		if m["logging.googleapis.com/spanId"] != sc.SpanID().String() || m["logging.googleapis.com/trace_sampled"] != true {
			t.Errorf("%s: spanId/trace_sampled = %v/%v", name, m["logging.googleapis.com/spanId"], m["logging.googleapis.com/trace_sampled"])
		}
		if th, _ := m["trace_header"].(map[string]any); th["trace_id"] != sc.TraceID().String() {
			t.Errorf("%s: trace_header = %v", name, m["trace_header"])
		}
		if m["function_name"] != "TestCtxMethods_AddTraceFieldsAndKeepCaller" {
			t.Errorf("%s: function_name = %v", name, m["function_name"])
		}
		if (name == "ErrorCtx" || name == "FatalCtx") && !strings.Contains(fmt.Sprint(m["error_cause"]), "TestCtxMethods_AddTraceFieldsAndKeepCaller") {
			t.Errorf("%s: error_cause = %v", name, m["error_cause"])
		}
	}
}

func TestPlainMethods_Unchanged(t *testing.T) {
	var buf bytes.Buffer
	l := newTestLogger(&buf)
	data := LogData{Description: "d", Err: errors.New("boom"), TraceHeader: map[string]string{"x": "y"}}

	l.Info(data)
	info := lastEntry(t, &buf)
	l.Error(data)
	errEntry := lastEntry(t, &buf)
	l.Alert(data)
	alert := lastEntry(t, &buf)
	l.Success(data)
	success := lastEntry(t, &buf)
	l.Fatal(data)
	fatal := lastEntry(t, &buf)

	entries := map[string]map[string]any{
		"Info": info, "Error": errEntry, "Alert": alert, "Success": success, "Fatal": fatal,
	}
	for name, m := range entries {
		for _, k := range []string{"trace_id", "span_id", "logging.googleapis.com/trace", "logging.googleapis.com/spanId"} {
			if _, ok := m[k]; ok {
				t.Errorf("%s: unexpected field %s", name, k)
			}
		}
		if th, _ := m["trace_header"].(map[string]any); th["x"] != "y" {
			t.Errorf("%s: trace_header = %v", name, m["trace_header"])
		}
		if m["function_name"] != "TestPlainMethods_Unchanged" {
			t.Errorf("%s: function_name = %v", name, m["function_name"])
		}
		if (name == "Error" || name == "Fatal") && !strings.Contains(fmt.Sprint(m["error_cause"]), "TestPlainMethods_Unchanged") {
			t.Errorf("%s: error_cause = %v", name, m["error_cause"])
		}
	}
}

func TestCtxMethods_NoSpanOrNilCtx(t *testing.T) {
	var buf bytes.Buffer
	l := newTestLogger(&buf)
	var nilCtx context.Context

	for name, ctx := range map[string]context.Context{"nil": nilCtx, "no span": context.Background()} {
		l.ErrorCtx(ctx, LogData{Description: "d", Err: errors.New("boom")})
		m := lastEntry(t, &buf)
		if _, ok := m["trace_id"]; ok {
			t.Errorf("%s ctx: unexpected trace_id", name)
		}
		if m["trace_header"] != nil {
			t.Errorf("%s ctx: trace_header = %v, want nil", name, m["trace_header"])
		}
	}
}

func TestAlert_ErrorCauseIsCaller(t *testing.T) {
	var buf bytes.Buffer
	l := newTestLogger(&buf)
	ctx, _ := spanCtx(t)

	l.Alert(LogData{Description: "d", Err: errors.New("boom")})
	plain := l.logData.errorCause
	l.AlertCtx(ctx, LogData{Description: "d", Err: errors.New("boom")})
	withCtx := l.logData.errorCause

	for name, cause := range map[string]string{"Alert": plain, "AlertCtx": withCtx} {
		if !strings.Contains(cause, "TestAlert_ErrorCauseIsCaller") {
			t.Errorf("%s: errorCause = %q, want caller frame", name, cause)
		}
	}
}
