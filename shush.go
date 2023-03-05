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
		return nil, fmt.Errorf("unexpected type: %T", pass.ResultOf[inspect.Analyzer])
	}
	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		sel := node.(*ast.SelectorExpr)
		pkgIdent, ok := sel.X.(*ast.Ident)
		if !ok {
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
