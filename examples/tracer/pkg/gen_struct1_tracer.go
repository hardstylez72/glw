package pkg

import (
	"net/http"

	"context"

	"math/big"

	"context"
	tracer "lwt/examples/tracer/tracerWrap"
)

type ServiceTracer struct {
	s Service
}

func (s *ServiceTracer) F5(h *http.Client) (err error) {

	_, span := tracer.Start(context.Background(), "pkg.F5")

	defer func() {
		span.End(err)
	}()

	return s.s.F5(h)

}

func (s *ServiceTracer) F1(g, f string, greg int) (err error) {

	_, span := tracer.Start(context.Background(), "pkg.F1")

	defer func() {
		span.End(err)
	}()

	return s.s.F1(g, f, greg)

}

func (s *ServiceTracer) F2(h bool) (err error) {

	return s.s.F2(h)

}

func (s *ServiceTracer) f3(ctx context.Context) (u context.Context) {

	ctx, span := tracer.Start(ctx, "pkg.f3")

	defer func() { span.End() }()

	return s.s.f3(ctx)

}

func (s *ServiceTracer) f4(a string, b, c int) (u *big.Int) {

	_, span := tracer.Start(context.Background(), "pkg.f4")

	defer func() { span.End() }()

	return s.s.f4(a, b, c)

}
