package model

type Trace struct {
	// populated in case it is from OTEL
	KustoOtelTraces []KustoOtelTrace
	// populated in case it is from jaeger spans
	KustoSpans []KustoSpan
	// type that indicates the source of the data, if this is from jaeger or otel
	sink string
}
