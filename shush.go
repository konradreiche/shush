// Package shush defines an Analyzer that reports usage of `fmt.Println`
// statements which may have been unintentionally left in the code for
// debugging purposes.
package shush

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `report fmt.Println statements

Find fmt.Println statements which may have been unintentionally left in the
code for debugging purposes.`

// Analyzer defines the analysis function to find and report fmt.Println
// statements. It can be referenced in an analysis driver.
var Analyzer = &analysis.Analyzer{
	Name:             "shush",
	Doc:              Doc,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	RunDespiteErrors: true,
	Run:              run,
}

func run(pass *analysis.Pass) (any, error) {
	inspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		// explicit check: should never be false due to inspect.Analyzer dependency
		return nil, fmt.Errorf("unexpected type: %T", pass.ResultOf[inspect.Analyzer])
	}
	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		sel := node.(*ast.SelectorExpr)
		pkgIdent, ok := sel.X.(*ast.Ident)
		if !ok {
			// explicit check: though nil map lookup below would catch this too
			return
		}
		pkgName, ok := pass.TypesInfo.Uses[pkgIdent].(*types.PkgName)
		if !ok || pkgName.Imported().Path() != "fmt" {
			return
		}
		if sel.Sel.Name == "Println" {
			pass.Reportf(sel.Pos(), "extraneous fmt.Println statement")
		}
	})
	return nil, nil
}
