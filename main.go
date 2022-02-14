package main

import (
	_ "embed"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

const (
	pathIsNotSet               = "package path is not set"
	idIsNotSet                 = "struct id is not set"
	loggerTemplateFileIdNotSet = "logger template file is not set"
	tracerTemplateFileIdNotSet = "tracer template file is not set"
	version                    = "0.0.1"
)

//go:embed templates/logger
var defaultLoggerTemplate string

//go:embed templates/interface
var defaultInterfaceTemplate string

//go:embed templates/tracer
var defaultTracerTemplate string

func main() {
	log.Println("glw version: " + version)
	var err error
	path := flag.String("path", pathIsNotSet, "/path/to/package")
	id := flag.String("id", idIsNotSet, "a string")
	force := flag.Bool("force", false, "deletes previous generated files")

	loggerTemplateFilePath := flag.String("logger_template_file", loggerTemplateFileIdNotSet, "see examples directory")
	logger := flag.Bool("logger", true, "do not generate logger file")

	tracerTemplateFilePath := flag.String("tracer_template_file", tracerTemplateFileIdNotSet, "see examples directory")
	tracer := flag.Bool("tracer", true, "do not generate tracer file")

	flag.Parse()

	if *path == pathIsNotSet {
		panic(pathIsNotSet)
	}

	if *id == idIsNotSet {
		panic(idIsNotSet)
	}

	var loggerTemplateFile string
	if *logger {
		if *loggerTemplateFilePath == loggerTemplateFileIdNotSet {
			loggerTemplateFile = defaultLoggerTemplate
			log.Println("you should specify your customized logger file template, using default one")
		} else {
			loggerTemplateFile, err = readFile(*loggerTemplateFilePath)
			if err != nil {
				panic(err)
			}
		}
	}

	var tracerTemplateFile string
	if *tracer {
		if *tracerTemplateFilePath == tracerTemplateFileIdNotSet {
			tracerTemplateFile = defaultTracerTemplate
			log.Println("you should specify your customized tracer file template, using default one")
		} else {
			tracerTemplateFile, err = readFile(*tracerTemplateFilePath)
			if err != nil {
				panic(err)
			}
		}
	}

	log.Println("glw generation started. path=" + *path + " id=" + *id)

	if err := Generate(*path, *id, *force, loggerTemplateFile, *logger, tracerTemplateFile, *tracer); err != nil {
		panic(err)
	}
	log.Println("glw generation finished. path=" + *path + " id=" + *id)
}

func readFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.New("can not open file: " + path)
	}
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return "", errors.New("can not read file: " + path)
	}
	return string(body), nil
}
