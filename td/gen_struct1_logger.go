package td

import (
	"net/http"

	"context"

	"math/big"
)

type serviceLogger struct {
	s Service
}

func (s *serviceLogger) F5(h *http.Client) (err error) {

	return s.s.F5(h)

}

func (s *serviceLogger) F1(g, f string, greg int) (err error) {

	return s.s.F1(g, f, greg)

}

func (s *serviceLogger) F2(h bool) (err error) {

	return s.s.F2(h)

}

func (s *serviceLogger) f3(ctx context.Context) (u context.Context) {

	return s.s.f3(ctx)

}

func (s *serviceLogger) f4(a string, b, c int) (u *big.Int) {

	return s.s.f4(a, b, c)

}
