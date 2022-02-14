package main

import "testing"

func TestGenerate(t *testing.T) {
	//t.Run("", func(t *testing.T) {
	//	Generate("/home/hs/Documents/bitsgap/match-engine/engine", "engine")
	//})

	t.Run("", func(t *testing.T) {
		Generate("examples/logger/pkg", "Struct1", false, defaultLoggerTemplate, true, defaultTracerTemplate, true)
	})

}
