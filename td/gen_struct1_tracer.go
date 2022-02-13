package td

import (
	"net/http"

	"context"

	"math/big"
)

type serviceTracer struct {
	s Service
}

func (s *serviceTracer) F5(h *http.Client) (err error) {

	return s.s.F5(h)

}

func (s *serviceTracer) F1(g, f string, greg int) (err error) {

	return s.s.F1(g, f, greg)

}

func (s *serviceTracer) F2(h bool) (err error) {

	return s.s.F2(h)

}

func (s *serviceTracer) f3(ctx context.Context) (u context.Context) {

	return s.s.f3(ctx)

}

func (s *serviceTracer) f4(a string, b, c int) (u *big.Int) {

	return s.s.f4(a, b, c)

}
