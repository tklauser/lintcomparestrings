// Copyright 2025 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/tklauser/lintcomparestring/comparestrings"
)

func main() {
	singlechecker.Main(comparestrings.Analyzer)
}
