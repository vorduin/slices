// Copyright © The Slices Author. All rights reserved.
// Use of this source logic, code, or documentation is governed by
// an MIT-style license that can be found in the LICENSE file.

package slices_test

import (
	"testing"
	"github.com/vorduin/slices"
)

func TestWithLen(t *testing.T) {
	s := slices.WithLen[int](10)

	if len(s) != 10 {
		t.Failed()
	}
}

func TestWithCap(t *testing.T) {
	s := slices.WithCap[int](10)

	if len(s) == 10 {
		t.Fail()
	}

	if cap(s) != 10 {
		t.Failed()
	}
}

func TestCopyInt(t *testing.T) {
	orig := []int{1, 2, 3}
	cp := slices.Copy(orig)
	orig[0] = 0

	if cp[0] == 0 {
		t.Failed()
	}
}

func TestCopyString(t *testing.T) {
	orig := []string{"the", "slices", "author"}
	cp := slices.Copy(orig)
	orig[0] = "a"

	if cp[0] == "a" {
		t.Failed()
	}
}

func TestCopy2DSlice(t *testing.T) {
	orig := [][]int{{1}, {2}, {3}}
	cp := slices.Copy(orig)
	orig[0] = []int{0}

	if cp[0][0] == 0 {
		t.Failed()
	}
}

func TestEqualDiffLen(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 4, 5}

	if slices.Equal(s1, s2) {
		t.Failed()
	}
}

func TestEqualInt(t *testing.T) {
	s1 := []int{1, 2, 3}
	if !slices.Equal(s1, s1) {
		t.Fail()
	}

	s2 := []int{0, 2, 3}
	if slices.Equal(s1, s2) {
		t.Fail()
	}
}

func TestEqualString(t *testing.T) {
	s1 := []string{"the", "slices", "author"}
	if !slices.Equal(s1, s1) {
		t.Fail()
	}

	s2 := []string{"a", "slices", "author"}
	if !slices.Equal(s1, s2) {
		t.Fail()
	}
}