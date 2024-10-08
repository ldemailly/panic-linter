package analyser

import (
	"go/ast"
	"strings"

	"fortio.org/sets"
	"golang.org/x/tools/go/analysis"
)

var Analyser = &analysis.Analyzer{
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
				if strings.HasPrefix(comment.Text, "/* want") { // skipping analysistests special comments
					continue
				}
				commentPos := pass.Fset.Position(comment.Pos())
				commentLines.Add(commentPos.Line)
			}
		}
		ast.Inspect(f, func(node ast.Node) bool {
			return CheckPanicCalls(pass, node, commentLines)
		})
	}
	return nil, nil //nolint:nilnil // we don't have a useful result to return.
}

type Result struct {
	IssueLines []int // Lines with issues, empty if no issues.
}

func CheckPanicCalls(pass *analysis.Pass, node ast.Node, commentLines sets.Set[int]) bool {
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
		pass.Reportf(node.Pos(), "panic call without same line comment justifying it")
	}
	return true
}
