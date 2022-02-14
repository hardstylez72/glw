package pkg

import (
	"net/http"

	"context"

	"math/big"
)

type Service interface {
	F5(h *http.Client) error

	F1(g, f string, greg int) error

	F2(h bool, m map[string]int, j chan int) error

	f3(ctx context.Context) context.Context

	f4(_ string, _, _ int) *big.Int
}
