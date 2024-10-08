package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"
)

func TestPanicCheck(t *testing.T) {
	src := `package main

import "fmt"

func main() {
	if false {
		fmt.Println("Nested Hello, world") // Comment inside if.
		panic("this is bad")
	} else {
		panic("this is ok") // A comment on the same line makes panic() ok.
	}
}
`
	fset := token.NewFileSet()
	// Parse the source code
	file, err := parser.ParseFile(fset, "example.go", src, parser.ParseComments)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	res := CheckPanicCalls(file, fset)
	if len(res.IssueLines) != 1 {
		t.Fatalf("Expected 1 issue, got %d", len(res.IssueLines))
	}
	if res.IssueLines[0] != 8 {
		t.Fatalf("Expected issue on line 8, got %d", res.IssueLines[0])
	}
}
