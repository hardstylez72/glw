package tracer

import (
	"context"
	"fmt"

	"go.opencensus.io/trace"
)

type span struct {
	source *trace.Span
	err    error
}

func FromContext(ctx context.Context) *span {
	return &span{source: trace.FromContext(ctx)}
}

func Start(ctx context.Context, name string) (context.Context, *span) {
	spanContext, tspan := trace.StartSpan(ctx, name)
	return spanContext, &span{source: tspan}
}
func (s *span) exist() bool {
	return s.source != nil
}

func (s *span) KV(key, value string) {
	if !s.exist() {
		return
	}
	s.source.Annotate([]trace.Attribute{trace.StringAttribute(key, value)}, "")
}

func (s *span) End(errs ...error) {
	if !s.exist() {
		return
	}
	for _, err := range errs {
		if err != nil {
			s.source.AddAttributes(trace.BoolAttribute("error", true))
			// если sprintF будет кушать много ресурсов, можно убрать
			s.source.Annotate([]trace.Attribute{trace.StringAttribute("stack", fmt.Sprintf("%+v", err))}, err.Error())
		}
	}

	s.source.End()
}
