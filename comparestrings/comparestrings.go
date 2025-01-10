// Copyright 2025 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package comparestrings defines an Analyzer that checks for calls to
// cmp.Compare that have string arguments and should use strings.Compare
// instead.
package comparestrings

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

const Doc = `check whether most efficient three-way-compare functions is used for strings

Using cmp.Compare to compare strings is less performant than strings.Compare,
especially for large strings since Go 1.23. See go.dev/issues/61725 and
go.dev/cl/532195 for details.`

var Analyzer = &analysis.Analyzer{
	Name:     "comparestring",
	Doc:      Doc,
	URL:      "https://github.com/tklauser/lintcomparestrings",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if !importsPackage(pass.Pkg, "cmp") {
		return nil, nil // doesn't directly import package cmp
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		fn, _ := typeutil.Callee(pass.TypesInfo, call).(*types.Func)
		if !isFunctionNamed(fn, "cmp", "Compare") {
			return
		}

		arg := call.Args[0]
		typ := pass.TypesInfo.Types[arg].Type

		if tuple, ok := typ.(*types.Tuple); ok {
			typ = tuple.At(0).Type() // special case for cmp.Compare(f(...), g(...))
		}

		b, ok := typ.Underlying().(*types.Basic)
		if !ok || b.Kind() != types.String {
			return
		}

		pass.Report(analysis.Diagnostic{
			Pos:     call.Pos(),
			End:     call.End(),
			Message: fmt.Sprintf("use strings.Compare instead of %s for three-way string comparison", fn.FullName()),
		})
	})

	return nil, nil
}

// importsPackage repoirts whether path is imported by pkg.
//
// Copied from
// golang.org/x/tools/go/analysis/passes/internal/analysisutil.Imports
func importsPackage(pkg *types.Package, path string) bool {
	for _, imp := range pkg.Imports() {
		if imp.Path() == path {
			return true
		}
	}
	return false
}

// isFunctionNamed reports whether f is a top-level function defined in the
// given package and has one of the given names.
// It returns false if f is nil or a method.
//
// Copied from
// golang.org/x/tools/go/analysis/passes/internal/analysisutil.IsFunctionNamed
func isFunctionNamed(f *types.Func, pkgPath string, names ...string) bool {
	if f == nil {
		return false
	}
	if f.Pkg() == nil || f.Pkg().Path() != pkgPath {
		return false
	}
	if f.Type().(*types.Signature).Recv() != nil {
		return false
	}
	for _, n := range names {
		if f.Name() == n {
			return true
		}
	}
	return false
}
