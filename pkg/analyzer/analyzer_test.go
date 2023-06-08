package analyzer_test

import (
	"testing"

	"github.com/adamdecaf/xmlencoderclose/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestImports(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "imports")
}

func TestGenerics(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "generics")
}

func TestMissing(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "missing")
}

func TestNoXML(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "noxml")
}

func TestValid(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.NewAnalyzer(), "valid")
}
