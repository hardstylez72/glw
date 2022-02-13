package main

import (
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Generate(path, id string, force bool, loggerTemplateFile string, generateLogger bool, tracerTemplateFile string, generateTracer bool) error {
	methods, pkg, deps, err := findStructAndMethods(path, id)
	if err != nil {
		return err
	}
	loggerFileName := makeLoggerFileName(id)
	tracerFileName := makeTracerFileName(id)
	interfaceFileName := makeInterfaceFileName(id)

	if force {
		_ = os.Remove(filepath.Join(path, interfaceFileName))

		if generateLogger {
			_ = os.Remove(filepath.Join(path, loggerFileName))
		}
		_ = os.Remove(filepath.Join(path, tracerFileName))
	}

	resolvedDeps := resolveDeps(pkg, deps)

	if err = generateInterfaceFile(methods, path, interfaceFileName, pkg.Name, defaultInterfaceTemplate, resolvedDeps); err != nil {
		log.Fatalln("can't generate interface file: " + path + "/" + interfaceFileName + " error: " + err.Error())
	}
	log.Println("interface file " + path + "/" + interfaceFileName + " generated")

	if generateLogger {
		if err = generateWrapFile(methods, path, loggerFileName, pkg.Name, loggerTemplateFile, resolvedDeps, "ServiceLogger"); err != nil {
			log.Fatalln("can't generate logger file: " + path + "/" + loggerTemplateFile + " error: " + err.Error())
		}
		log.Println("logger file " + path + "/" + loggerFileName + " generated")
	}

	if generateTracer {
		if err = generateWrapFile(methods, path, tracerFileName, pkg.Name, tracerTemplateFile, resolvedDeps, "ServiceTracer"); err != nil {
			log.Fatalln("can't generate tracer file: " + path + "/" + tracerTemplateFile + " error: " + err.Error())
		}
		log.Println("tracer file " + path + "/" + tracerFileName + " generated")
	}

	return nil
}

func makeLoggerFileName(id string) string {
	return "gen_" + strings.ToLower(id) + "_logger.go"
}

func makeInterfaceFileName(id string) string {
	return "gen_" + strings.ToLower(id) + "_interface.go"
}

func makeTracerFileName(id string) string {
	return "gen_" + strings.ToLower(id) + "_tracer.go"
}

func resolveDeps(pkg *build.Package, deps []string) []string {
	result := make([]string, 0)
	for _, d := range deps {
		i := index(pkg.Imports, d)
		if i != -1 {
			result = append(result, pkg.Imports[i])
		}
	}

	m := make(map[string]bool)

	for _, el := range result {
		m[el] = true
	}
	result = make([]string, 0)
	for key := range m {
		result = append(result, key)
	}

	return result
}

func index(ref []string, target string) int {
	for i, el := range ref {
		slash := strings.LastIndex(el, "/")
		if slash == -1 {

			if el == target {
				return i
			}
		} else {
			if el[slash+1:] == target {
				return i
			}
		}

	}
	return -1
}
