package model

import (
	"time"

	"github.com/Azure/azure-kusto-go/kusto/data/value"
)

type KustoOtelTrace struct {
	TraceID            string                 `kusto:"TraceID"`            // TraceID associated to the Trace
	SpanID             string                 `kusto:"SpanID"`             // SpanID associated to the Trace
	ParentID           string                 `kusto:"ParentID"`           // ParentID associated to the Trace
	SpanName           string                 `kusto:"SpanName"`           // The SpanName of the Trace
	SpanStatus         string                 `kusto:"SpanStatus"`         // The SpanStatus associated to the Trace
	SpanKind           string                 `kusto:"SpanKind"`           // The SpanKind of the Trace
	StartTime          string                 `kusto:"StartTime"`          // The start time of the occurrence. Formatted into string as RFC3339Nano
	EndTime            string                 `kusto:"EndTime"`            // The end time of the occurrence. Formatted into string as RFC3339Nano
	ResourceAttributes map[string]interface{} `kusto:"ResourceAttributes"` // JSON Resource attributes that can then be parsed.
	TraceAttributes    map[string]interface{} `kusto:"TraceAttributes"`    // JSON attributes that can then be parsed.
	Events             []*OtelEvent           `kusto:"Events"`             // Array containing the events in a span
	Links              []*OtelLink            `kusto:"Links"`              // Array containing the link in a span
}

type OtelEvent struct {
	EventName       string                 `kusto:"EventName"`       // Array containing the events in a span
	Timestamp       string                 `kusto:"Timestamp"`       // Array containing the events in a span
	EventAttributes map[string]interface{} `kusto:"EventAttributes"` // Array containing the events in a span
}

// Follows from links
type OtelLink struct {
	TraceID            string                 `kusto:"TraceID"`
	SpanID             string                 `kusto:"SpanID"`
	TraceState         string                 `kusto:"TraceState"`
	SpanLinkAttributes map[string]interface{} `kusto:"SpanLinkAttributes"`
}

type KustoSpan struct {
	TraceID            string        `kusto:"TraceID"`
	SpanID             string        `kusto:"SpanID"`
	OperationName      string        `kusto:"OperationName"`
	References         value.Dynamic `kusto:"References"`
	Flags              int32         `kusto:"Flags"`
	StartTime          time.Time     `kusto:"StartTime"`
	Duration           time.Duration `kusto:"Duration"`
	Tags               value.Dynamic `kusto:"Tags"`
	Logs               value.Dynamic `kusto:"Logs"`
	ProcessServiceName string        `kusto:"ProcessServiceName"`
	ProcessTags        value.Dynamic `kusto:"ProcessTags"`
	ProcessID          string        `kusto:"ProcessID"`
}
