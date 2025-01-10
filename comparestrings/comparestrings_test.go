// Copyright 2025 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comparestrings_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/tklauser/lintcomparestring/comparestrings"
)

func TestAllAnalysis(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), comparestrings.Analyzer)
}
