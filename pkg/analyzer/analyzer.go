package analyzer

import (
	"errors"
	"go/token"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "xmlencoderclose",
		Doc:  "Checks that xml.Encoder is closed.",
		Run:  run,
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs //nolint:forcetypeassert

	// Quit early of the file doesn't use xml.Encoder
	if !imports(pass.Pkg, "encoding/xml") {
		return nil, nil
	}

	encType := analysisutil.TypeOf(pass, "encoding/xml", "*Encoder")
	if encType == nil {
		return nil, errors.New("analyzer did not find xml.Encoder type")
	}

	var methods []*types.Func
	if m := analysisutil.MethodOf(encType, "Close"); m != nil {
		methods = append(methods, m)
	}
	if len(methods) == 0 {
		return nil, errors.New("no method finders")
	}

	for _, f := range funcs {
		for _, b := range f.Blocks {
			for i := range b.Instrs {
				var pos token.Pos
				switch instr := b.Instrs[i].(type) {
				case *ssa.Extract:
					pos = instr.Tuple.Pos()
				default:
					pos = instr.Pos()
				}
				called, ok := analysisutil.CalledFrom(b, i, encType, methods...)
				if ok && !called {
					pass.Reportf(pos, "Encoder.Close must be called")
				}
			}
		}
	}

	return nil, nil
}

func imports(pkg *types.Package, path string) bool {
	for _, imp := range pkg.Imports() {
		if imp.Path() == path {
			return true
		}
	}
	return false
}
