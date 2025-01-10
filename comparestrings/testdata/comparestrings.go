// Copyright 2025 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import (
	"cmp"
	"fmt"
	"strings"
)

type t struct{}

func (tp *t) String() string {
	return "foo"
}

func adjust(s string) string {
	return fmt.Sprintf("<%s>", s)
}

func CmpCompare() {
	var x, y int
	x, y = 42, 23
	_ = cmp.Compare(0, 1)
	_ = cmp.Compare(x, y)

	a, b := "A", "B"
	aa, bb := &t{}, &t{}
	_ = cmp.Compare(a, b)                     // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(aa.String(), bb.String()) // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(adjust(a), adjust(b))     // want `use strings.Compare instead of cmp.Compare for three-way string comparison`

	_ = strings.Compare(a, b)
	_ = strings.Compare(aa.String(), bb.String())
	_ = strings.Compare(adjust(a), adjust(b))
}
