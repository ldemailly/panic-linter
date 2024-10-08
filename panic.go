package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"

	"fortio.org/sets"
)

func main() {
	singlechecker.Main(analyzer)
}

var analyzer = &analysis.Analyzer{
	Name: "paniccheck",
	Doc:  "reports panic() calls without comments",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		// Create a map of lines that have comments
		commentLines := sets.New[int]()
		for _, commentGroup := range f.Comments {
			for _, comment := range commentGroup.List {
				commentPos := pass.Fset.Position(comment.Pos())
				commentLines.Add(commentPos.Line)
			}
		}
		ast.Inspect(f, func(node ast.Node) bool {
			return CheckPanicCalls(pass, f, node, commentLines)
		})
	}
	return nil, nil
}

type Result struct {
	IssueLines []int // Lines with issues, empty if no issues.
}

func CheckPanicCalls(pass *analysis.Pass, file *ast.File, node ast.Node, commentLines sets.Set[int]) bool {
	// Check for a panic call
	callExpr, ok := node.(*ast.CallExpr)
	if !ok {
		return true
	}
	funIdent, ok := callExpr.Fun.(*ast.Ident)
	if !ok || funIdent.Name != "panic" {
		return true
	}

	startPos := pass.Fset.Position(node.Pos())
	// Check if the line has a comment
	if !commentLines.Has(startPos.Line) {
		pass.Reportf(node.Pos(), "panic without comment")
	}
	return true
}
