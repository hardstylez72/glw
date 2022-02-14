package logger

import (
	"github.com/hardstylez72/glw/examples/logger/pkg"
	"net/http"
)

func main() {
	var s pkg.Service
	s = &pkg.Struct1{}
	s = &pkg.ServiceLogger{Service: s}
	s = &pkg.ServiceTracer{Service: s}
	_ = s.F5(&http.Client{})
}
