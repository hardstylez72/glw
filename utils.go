package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"strings"
)

type Param struct {
	Names []string
	Type  string
}

func getParamDeps(f *ast.FuncDecl) []string {

	result := make([]string, 0)

	if f.Type.Params != nil {
		for _, p := range f.Type.Params.List {

			if star, ok := p.Type.(*ast.SelectorExpr); ok {
				if i, ok := star.X.(*ast.Ident); ok {
					if i.Name != "" {
						result = append(result, i.Name)
					}
				}
			}

			if star, ok := p.Type.(*ast.StarExpr); ok {

				if star, ok := star.X.(*ast.StarExpr); ok {
					if i, ok := star.X.(*ast.Ident); ok {
						if i.Name != "" {
							result = append(result, i.Name)
						}
					}
				}

				if star, ok := star.X.(*ast.SelectorExpr); ok {
					if i, ok := star.X.(*ast.Ident); ok {
						if i.Name != "" {
							result = append(result, i.Name)
						}
					}
				}

			}
		}
	}

	if f.Type.Results != nil {
		for _, p := range f.Type.Results.List {

			if star, ok := p.Type.(*ast.SelectorExpr); ok {
				if i, ok := star.X.(*ast.Ident); ok {
					if i.Name != "" {
						result = append(result, i.Name)
					}
				}
			}

			if star, ok := p.Type.(*ast.StarExpr); ok {

				if star, ok := star.X.(*ast.StarExpr); ok {
					if i, ok := star.X.(*ast.Ident); ok {
						if i.Name != "" {
							result = append(result, i.Name)
						}
					}
				}

				if star, ok := star.X.(*ast.SelectorExpr); ok {
					if i, ok := star.X.(*ast.Ident); ok {
						if i.Name != "" {
							result = append(result, i.Name)
						}
					}
				}
			}
		}
	}

	return result
}

func printFuncParamsNames(params []string) string {
	return strings.Join(params, ", ")
}

func normalizeInputFuncParams(method Method) Method {

	runeName := 'a'

	for i := range method.Params {
		for j := range method.Params[i].Names {
			if method.Params[i].Type == "context.Context" {
				method.Params[i].Names[j] = NormalizedInputContextName
				method.MethodOptions.ContextInputToken = NormalizedInputContextName
				method.MethodOptions.HasContext = true
			}

			if method.Params[i].Names[j] == "_" {
				method.Params[i].Names[j] = string(runeName)
				runeName++
			}
		}
	}
	return method
}

func normalizeOutputFuncParams(method Method) Method {

	runeName := 'u'

	for i := range method.Returns {
		if method.Returns[i].Type == "error" {
			method.Returns[i].Names = []string{NormalizedOutputErrorName}
			method.MethodOptions.ErrorOutputToken = NormalizedOutputErrorName
			method.MethodOptions.HasError = true
		} else {
			names := make([]string, 0)

			if len(method.Returns[i].Names) == 0 {
				names = append(names, string(runeName))
				runeName++
			} else {
				for range method.Returns[i].Names {
					names = append(names, string(runeName))
					runeName++
				}
			}

			method.Returns[i].Names = names
		}
	}
	return method
}

// gofmt pretty-prints e.
func gofmt(e ast.Expr) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, fileSet, e)
	return buf.String()
}
