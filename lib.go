package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
)

var fileSet = token.NewFileSet()

func findStructAndMethods(path, id string) ([]Method, *build.Package, []string, error) {

	pkg, err := findPkg(path)
	if err != nil {
		return nil, nil, nil, err
	}

	files, err := parsePkgFiles(pkg, fileSet)
	if err != nil {
		return nil, nil, nil, err
	}

	targetSpec, err := findStructTypeSpec(files, id)
	if err != nil {
		return nil, nil, nil, err
	}

	methods, err := findStructMethods(files, targetSpec)
	if err != nil {
		return nil, nil, nil, err
	}

	deps := make([]string, 0)
	result := make([]Method, 0)
	for _, method := range methods {
		result = append(result, toMethod(method))
		deps = append(deps, getParamDeps(method)...)
	}

	return result, pkg, deps, nil
}

func parsePkgFiles(pkg *build.Package, fileSet *token.FileSet) ([]*ast.File, error) {
	files := pkg.GoFiles
	result := make([]*ast.File, 0)

	for _, file := range files {
		f, err := parser.ParseFile(fileSet, filepath.Join(pkg.Dir, file), nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}

func findStructMethods(files []*ast.File, target *ast.TypeSpec) ([]*ast.FuncDecl, error) {

	result := make([]*ast.FuncDecl, 0)

	for _, f := range files {

		for _, decl := range f.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if ok {
				if funcDecl.Recv != nil {
					for _, field := range funcDecl.Recv.List {

						// pointer to the target
						if star, ok := field.Type.(*ast.StarExpr); ok {
							if i, ok := star.X.(*ast.Ident); ok {
								if i.Obj != nil {
									if i.Obj.Decl == target {
										result = append(result, funcDecl)
									}
								} else {
									if i.Name == target.Name.Name {
										result = append(result, funcDecl)
									}
								}
							}
						}
						// direct struct method
						if i, ok := field.Type.(*ast.Ident); ok {
							if i.Obj != nil {
								if i.Obj.Decl == target {
									result = append(result, funcDecl)
								}
							} else {
								if i.Name == target.Name.Name {
									result = append(result, funcDecl)
								}
							}

						}
					}
				}
			}
		}
	}
	return result, nil
}

func findStructTypeSpec(files []*ast.File, id string) (*ast.TypeSpec, error) {

	var structTypeSpec *ast.TypeSpec

	for _, f := range files {

		for _, decl := range f.Decls {
			target, ok := decl.(*ast.GenDecl)
			if ok {
				for _, spec := range target.Specs {
					targetSpec, ok := spec.(*ast.TypeSpec)
					if ok {
						if targetSpec.Name.Name == id {
							obj := targetSpec.Name.Obj
							if obj.Kind == ast.Typ {
								structTypeSpec = targetSpec
							}
						}
					}
				}
			}
		}
	}

	if structTypeSpec == nil {
		return nil, errors.New("struct with id: " + id + " is not found")
	}

	return structTypeSpec, nil
}

func findPkg(path string) (*build.Package, error) {
	pkg1, err := build.ImportDir(path, 0)
	if err == nil {
		return pkg1, nil
	}
	pkg, err := build.Import(path, ".", 0)
	if err != nil {
		return nil, fmt.Errorf("couldn't find package %s: %v", path, err)
	}
	return pkg, nil
}
