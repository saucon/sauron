package log

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

var (
	// gcpProjectID is needed to build the Cloud Logging trace resource name.
	// It stays empty off GCP, and until the background lookup finishes; logs
	// written before that still carry trace_id/span_id, just not the link.
	gcpProjectID       atomic.Value // string
	resolveProjectOnce sync.Once
)

// withTrace returns data (with TraceHeader filled if unset) and the log
// fields for the active span in ctx. With no valid span it returns data
// unchanged and nil fields, so the log line matches the plain method.
func withTrace(ctx context.Context, data LogData) (LogData, logrus.Fields) {
	if ctx == nil {
		return data, nil
	}
	sc := trace.SpanContextFromContext(ctx)
	if !sc.IsValid() {
		return data, nil
	}
	resolveGCPProjectID()

	traceID, spanID := sc.TraceID().String(), sc.SpanID().String()
	if data.TraceHeader == nil {
		data.TraceHeader = map[string]string{"trace_id": traceID, "span_id": spanID}
	}

	// logging.googleapis.com/* keys let Cloud Logging link the entry to
	// Cloud Trace: https://cloud.google.com/logging/docs/structured-logging
	fields := logrus.Fields{
		"trace_id":                             traceID,
		"span_id":                              spanID,
		"logging.googleapis.com/spanId":        spanID,
		"logging.googleapis.com/trace_sampled": sc.IsSampled(),
	}
	if projectID, _ := gcpProjectID.Load().(string); projectID != "" {
		fields["logging.googleapis.com/trace"] = "projects/" + projectID + "/traces/" + traceID
	}
	return data, fields
}

// resolveGCPProjectID looks the project ID up once, off the logging path:
// GOOGLE_CLOUD_PROJECT if set, else the metadata server in the background so
// a log call never blocks on the network.
func resolveGCPProjectID() {
	resolveProjectOnce.Do(func() {
		if id := os.Getenv("GOOGLE_CLOUD_PROJECT"); id != "" {
			gcpProjectID.Store(id)
			return
		}
		go func() {
			if !metadata.OnGCE() {
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if id, err := metadata.ProjectIDWithContext(ctx); err == nil {
				gcpProjectID.Store(id)
			}
		}()
	})
}
