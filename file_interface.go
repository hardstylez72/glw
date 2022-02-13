package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

type InterfaceFile struct {
	Methods []MethodPrint
	PkgName string
	Imports []string
}

func PrepareInterfaceFileStruct(methods []Method, pkgName string, deps []string) InterfaceFile {
	ms := make([]MethodPrint, 0)
	for _, m := range methods {

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

		ms = append(ms, MethodPrint{
			MethodName:    m.Name,
			MethodParams:  params,
			MethodReturns: returns,
		})
	}
	return InterfaceFile{
		Methods: ms,
		PkgName: pkgName,
		Imports: deps,
	}
}

func generateInterfaceFile(methods []Method, path, name, pkgName, templateBody string, deps []string) error {

	t, err := template.New("interfaceFileTemplate").Parse(templateBody)
	if err != nil {
		return err
	}

	sourceTemplate := PrepareInterfaceFileStruct(methods, pkgName, deps)

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
