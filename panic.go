package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"fortio.org/sets"
)

func main() {
	fset := token.NewFileSet()
	for _, filePath := range os.Args[1:] {
		if filePath == "--" { // to be able to run this like "go run main.go -- input.go"
			continue
		}
		f, err := parser.ParseFile(fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("Failed to parse file %s: %s", filePath, err)
		}
		CheckPanicCalls(f, fset)
	}
}

type Result struct {
	IssueLines []int // Lines with issues, empty if no issues.
}

func CheckPanicCalls(file *ast.File, fset *token.FileSet) (res Result) {
	// Create a map of lines that have comments
	commentLines := sets.New[int]()
	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			commentPos := fset.Position(comment.Pos())
			commentLines.Add(commentPos.Line)
		}
	}

	// Walk the AST and check for panic calls and commented lines
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		// Check for a panic call
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if funIdent, ok := callExpr.Fun.(*ast.Ident); ok && funIdent.Name == "panic" {
				startPos := fset.Position(n.Pos())
				// Check if the line has a comment
				if !commentLines.Has(startPos.Line) {
					// How to get back the full original line?
					fmt.Fprintf(os.Stderr, "%s:%d: panic() call without comment\n",
						file.Name, startPos.Line)
					res.IssueLines = append(res.IssueLines, startPos.Line)
				}
			}
		}
		return true
	})
	return
}

/*
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[1:] {
		if filePath == "--" { // to be able to run this like "go run main.go -- input.go"
			continue
		}
		f, err := parser.ParseFile(v.fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("Failed to parse file %s: %s", filePath, err)
		}

		ast.Walk(&v, f)
	}
}

type visitor struct {
	fset *token.FileSet
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, v.fset, node)
	fmt.Printf("%s | comment %d\n", buf.String(), len(node.Comments))

	return v
}

// test of what we want to flag

func ShouldBeFlagged() {
	panic("this should be flagged")
}

func ShouldNotBeFlagged() {
	panic("this should not be flagged") // thanks to the comment
}
*/
