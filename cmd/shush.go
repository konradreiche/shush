package main

import (
	"github.com/konradreiche/shush"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(shush.Analyzer)
}
