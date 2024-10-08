package main

import (
	"github.com/ldemailly/panic-linter/analyser"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyser.Analyser)
}
