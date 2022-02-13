package pkg

import (
	"context"
	"math/big"
)

//go:generate glw -path=. -id=Struct1 -logger_template_file=../logger_example -tracer=false
type Struct1 struct {
}

func (s *Struct1) F1(g, f string, greg int) error {
	return nil
}

// glw-no-logger
func (s *Struct1) F2(h bool) error {

	return nil
}

func (s *Struct1) f3(ctx context.Context) context.Context {
	return ctx
}

func (s Struct1) f4(_ string, _, _ int) *big.Int {
	return nil
}

type fff struct {
}

func (s fff) M1() error {
	return nil
}

func m() error {
	return nil
}
