package pkg

import (
	"context"

	"math/big"

	"net/http"

	"log"
)

type ServiceLogger struct {
	Service Service
}

func (s *ServiceLogger) F5(h *http.Client) (err error) {

	log.Println("pkg.F5 start")

	defer func() {
		if err != nil {
			log.Println("pkg.F5 error " + err.Error())
		} else {
			log.Println("pkg.F5 finish")
		}
	}()

	return s.Service.F5(h)

}

func (s *ServiceLogger) F1(g, f string, greg int) (err error) {

	log.Println("pkg.F1 start")

	defer func() {
		if err != nil {
			log.Println("pkg.F1 error " + err.Error())
		} else {
			log.Println("pkg.F1 finish")
		}
	}()

	return s.Service.F1(g, f, greg)

}

func (s *ServiceLogger) F2(h bool) (err error) {

	return s.Service.F2(h)

}

func (s *ServiceLogger) f3(ctx context.Context) (u context.Context) {

	log.Println("pkg.f3 start")

	defer func() {
		log.Println("pkg.f3 finish")
	}()

	return s.Service.f3(ctx)

}

func (s *ServiceLogger) f4(a string, b, c int) (u *big.Int) {

	log.Println("pkg.f4 start")

	defer func() {
		log.Println("pkg.f4 finish")
	}()

	return s.Service.f4(a, b, c)

}
