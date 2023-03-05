package shush

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/atomic"
)

func TestShush(t *testing.T) {
	path := analysistest.TestData()
	analysistest.Run(t, path, atomic.Analyzer, "a")
}
