package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	_ = analysistest.Run(t, "./testdata", analyzer, "./...")

}
