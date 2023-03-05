package shush

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/packages"
)

func TestShush(t *testing.T) {
	t.Run("test-data", func(t *testing.T) {
		path := analysistest.TestData()
		analysistest.Run(t, path, Analyzer, "a")
	})

	t.Run("self", func(t *testing.T) {
		dir, err := filepath.Abs(".")
		if err != nil {
			t.Fatal(err)
		}
		analysistest.Run(t, dir, Analyzer, "shush")
	})

}

func TestRequiredAnalyzer(t *testing.T) {
	mode := packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports |
		packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo |
		packages.NeedDeps

	cfg := &packages.Config{
		Mode: mode,
		Dir:  analysistest.TestData(),
	}

	pkgs, err := packages.Load(cfg, "a")
	if err != nil {
		t.Fatal(err)
	}
	for _, pkg := range pkgs {
		pass := &analysis.Pass{
			Analyzer:     Analyzer,
			Fset:         pkg.Fset,
			Files:        pkg.Syntax,
			OtherFiles:   pkg.OtherFiles,
			IgnoredFiles: pkg.IgnoredFiles,
			Pkg:          pkg.Types,
			TypesInfo:    pkg.TypesInfo,
			TypesSizes:   pkg.TypesSizes,
			TypeErrors:   pkg.TypeErrors,
			Report:       func(d analysis.Diagnostic) {},
		}
		if _, err := run(pass); err == nil {
			t.Fatal("expected run to fail")
		}
	}
}
