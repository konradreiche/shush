package shush

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestShush(t *testing.T) {
	path := analysistest.TestData()
	analysistest.Run(t, path, Analyzer, "a")
}
