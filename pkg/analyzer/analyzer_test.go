package analyzer_test

import (
	"testing"

	"github.com/adamdecaf/xmlencoderclose/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

var testAnalyzer = analyzer.NewAnalyzer()

func TestImports(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer, "imports")
}

func TestGenerics(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer, "generics")
}

func TestMissing(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer, "missing")
}

func TestNoXML(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer, "noxml")
}

func TestValid(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer, "valid")
}
