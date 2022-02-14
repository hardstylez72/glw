package pkg

import (
	"net/http"

	"context"

	"math/big"
)

type ServiceTracer struct {
	Service Service
}

func (s *ServiceTracer) F5(h *http.Client) (err error) {

	return s.Service.F5(h)

}

func (s *ServiceTracer) F1(g, f string, greg int) (err error) {

	return s.Service.F1(g, f, greg)

}

func (s *ServiceTracer) F2(h bool, m map[string]int, j chan int) (err error) {

	return s.Service.F2(h, m, j)

}

func (s *ServiceTracer) f3(ctx context.Context) (u context.Context) {

	return s.Service.f3(ctx)

}

func (s *ServiceTracer) f4(a string, b, c int) (u *big.Int) {

	return s.Service.f4(a, b, c)

}
