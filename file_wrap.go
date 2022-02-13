package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

type WrapFile struct {
	Methods    []MethodPrint
	StructName string
	PkgName    string
	Imports    []string
}

type MethodPrint struct {
	MethodName        string
	MethodParams      []ParamPrint
	MethodReturns     []ReturnPrint
	MethodRecv        string
	MethodStructName  string
	WithLogger        bool
	WithTracer        bool
	HasError          bool
	ErrorOutputToken  string
	HasContext        bool
	ContextInputToken string

	PkgName string
}

type ParamPrint struct {
	ParamName string
	ParamType string
}

type ReturnPrint struct {
	ReturnName string
	ReturnType string
}

func PrepareWrapFileStruct(methods []Method, pkgName string, deps []string, structName string) WrapFile {

	methodsPrint := make([]MethodPrint, 0)
	for _, m := range methods {

		m = normalizeInputFuncParams(m)
		m = normalizeOutputFuncParams(m)

		params := make([]ParamPrint, 0)
		for _, p := range m.Params {
			params = append(params, ParamPrint{
				ParamName: printFuncParamsNames(p.Names),
				ParamType: p.Type,
			})
		}

		returns := make([]ReturnPrint, 0)
		for _, p := range m.Returns {
			returns = append(returns, ReturnPrint{
				ReturnName: printFuncParamsNames(p.Names),
				ReturnType: p.Type,
			})
		}
		methodsPrint = append(methodsPrint, MethodPrint{
			MethodName:        m.Name,
			MethodParams:      params,
			MethodReturns:     returns,
			MethodRecv:        "s",
			MethodStructName:  structName,
			WithLogger:        m.MethodOptions.WithLogger,
			WithTracer:        m.MethodOptions.WithTracer,
			HasError:          m.MethodOptions.HasError,
			ErrorOutputToken:  m.MethodOptions.ErrorOutputToken,
			HasContext:        m.MethodOptions.HasContext,
			ContextInputToken: m.MethodOptions.ContextInputToken,
			PkgName:           pkgName,
		})
	}

	return WrapFile{
		Imports:    deps,
		Methods:    methodsPrint,
		StructName: structName,
		PkgName:    pkgName,
	}
}

func generateWrapFile(methods []Method, path, name, pkgName, templateBody string, deps []string, structName string) error {

	t, err := template.New("loggerFileTemplate").Parse(templateBody)
	if err != nil {
		return err
	}

	sourceTemplate := PrepareWrapFileStruct(methods, pkgName, deps, structName)

	buf := bytes.NewBuffer([]byte{})

	err = t.Execute(buf, &sourceTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path, name))
	if err != nil {
		return err
	}

	content, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}
