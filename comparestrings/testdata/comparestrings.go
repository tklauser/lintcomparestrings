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

type s string

type sa = string

type saa = s

func CmpCompare() {
	var x, y int
	x, y = 42, 23
	_ = cmp.Compare(0, 1)
	_ = cmp.Compare(x, y)

	a1, b1 := "A", "B"
	a2, b2 := &t{}, &t{}
	a3, b3 := s("aaa"), s("bbb")
	a4, b4 := sa(a1), sa(b1)
	a5, b5 := saa(a4), saa(b4)
	_ = cmp.Compare(a1, b1)                   // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(a2.String(), b2.String()) // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(adjust(a1), adjust(b1))   // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(a3, b3)                   // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(a4, b4)                   // want `use strings.Compare instead of cmp.Compare for three-way string comparison`
	_ = cmp.Compare(a5, b5)                   // want `use strings.Compare instead of cmp.Compare for three-way string comparison`

	_ = strings.Compare(a1, b1)
	_ = strings.Compare(a2.String(), b2.String())
	_ = strings.Compare(adjust(a1), adjust(b1))
}
