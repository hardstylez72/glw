package main

import (
	"go/ast"
	"strings"
)

const (
	MethodOptionCommentNoLogger = "glw-no-logger"
	MethodOptionCommentNoTracer = "glw-no-tracer"
	NormalizedOutputErrorName   = "err"
	NormalizedInputContextName  = "ctx"
)

type Method struct {
	Doc           []string
	Name          string
	Params        []Param
	Returns       []Param
	MethodOptions MethodOptions
}

type MethodOptions struct {
	WithLogger bool
	WithTracer bool

	HasContext        bool
	ContextInputToken string

	HasError         bool
	ErrorOutputToken string
}

func toMethod(f *ast.FuncDecl) Method {

	params := make([]Param, 0)
	if f.Type.Params != nil {
		for _, param := range f.Type.Params.List {

			names := make([]string, 0)
			for _, name := range param.Names {
				names = append(names, name.Name)
			}

			params = append(params, Param{
				Names: names,
				Type:  gofmt(param.Type),
			})
		}
	}

	returns := make([]Param, 0)
	if f.Type.Results != nil {
		for _, param := range f.Type.Results.List {

			names := make([]string, 0)
			for _, name := range param.Names {
				names = append(names, name.Name)
			}

			returns = append(returns, Param{
				Names: names,
				Type:  gofmt(param.Type),
			})
		}
	}

	doc := make([]string, 0)

	if f.Doc != nil {
		for _, line := range f.Doc.List {
			doc = append(doc, line.Text)
		}
	}

	return Method{
		Name:          f.Name.Name,
		Params:        params,
		Returns:       returns,
		Doc:           doc,
		MethodOptions: getMethodOptions(doc),
	}
}

func getMethodOptions(comments []string) MethodOptions {
	opt := MethodOptions{
		WithLogger: true,
		WithTracer: true,
	}

	comment := strings.Join(comments, " ")

	if strings.Contains(comment, MethodOptionCommentNoLogger) {
		opt.WithLogger = false
	}

	if strings.Contains(comment, MethodOptionCommentNoTracer) {
		opt.WithTracer = false
	}
	return opt
}
