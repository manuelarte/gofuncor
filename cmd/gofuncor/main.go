package main

import (
	"github.com/manuelarte/gofuncor/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}

func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.NewAnalyzer()}, nil
}
