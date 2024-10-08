package main

import (
	"testing"

	"github.com/ldemailly/panic-linter/analyser"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	_ = analysistest.Run(t, "./testdata", analyser.Analyser, "./...")

}
