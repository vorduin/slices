// Copyright © The Slices Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"testing"

	"github.com/vorduin/slices"
)

func TestWithLen(t *testing.T) {
	s := slices.WithLen[int](10)

	if len(s) != 10 {
		t.Fail()
	}

	s = slices.WithLen[int](0)
	if s == nil {
		t.Fail()
	}
}

func TestWithCap(t *testing.T) {
	s := slices.WithCap[int](10)

	if len(s) == 10 {
		t.Fail()
	}

	if cap(s) != 10 {
		t.Fail()
	}

	s = slices.WithCap[int](0)
	if s == nil {
		t.Fail()
	}
}

func TestClone(t *testing.T) {
	orig := []int{1, 2, 3}
	cp := slices.Clone(orig)
	orig[0] = 0

	if cp[0] == 0 {
		t.Fail()
	}

	orig = []int{}
	cp = slices.Clone(orig)
	if cp == nil {
		t.Fail()
	}

	orig = nil
	cp = slices.Clone(orig)
	if cp != nil {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 2, 3}
	rev := slices.Reverse(s)

	for i := 0; i < len(s); i++ {
		if s[i] != rev[len(rev)-1-i] {
			t.Fail()
		}
	}

	s = nil
	rev = slices.Reverse(s)
	if rev != nil {
		t.Fail()
	}
}

func TestLast(t *testing.T) {
	s := []int{1, 2, 3}

	if slices.Last(s) != s[len(s)-1] {
		t.Failed()
	}
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3}

	if !slices.Contains(s, 2) {
		t.Fail()
	}

	if slices.Contains(s, 4) {
		t.Fail()
	}

	s = nil
	if slices.Contains(s, 1) {
		t.Fail()
	}
}

func TestCount(t *testing.T) {
	s := []int{1, 2, 2, 3}

	if slices.Count(s, 0) != 0 {
		t.Fail()
	}

	if slices.Count(s, 1) != 1 {
		t.Fail()
	}

	if slices.Count(s, 2) != 2 {
		t.Fail()
	}

	s = nil
	if slices.Count(s, 1) != 0 {
		t.Fail()
	}
}

func TestIndex(t *testing.T) {
	s := []int{1, 2, 2, 3}

	if slices.Index(s, 2) != 1 {
		t.Fail()
	}

	if slices.Index(s, 0) != -1 {
		t.Fail()
	}

	s = nil
	if slices.Index(s, 1) != -1 {
		t.Fail()
	}
}

func TestJoin(t *testing.T) {
	lhs := []int{1, 2, 3}
	rhs := []int{4, 5, 6}
	j := slices.Join(lhs, rhs)
	if !slices.Equal(j, []int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}

	lhs = []int{}
	j = slices.Join(lhs, rhs)
	if !slices.Equal(j, []int{4, 5, 6}) {
		t.Fail()
	}

	lhs = nil
	j = slices.Join(lhs, rhs)
	if !slices.Equal(j, []int{4, 5, 6}) {
		t.Fail()
	}

	lhs = []int{1, 2, 3}
	rhs = []int{}
	j = slices.Join(lhs, rhs)
	if !slices.Equal(j, []int{1, 2, 3}) {
		t.Fail()
	}

	rhs = nil
	j = slices.Join(lhs, rhs)
	if !slices.Equal(j, []int{1, 2, 3}) {
		t.Fail()
	}

	lhs = []int{}
	rhs = []int{}
	j = slices.Join(lhs, rhs)
	if j == nil || len(j) != 0 {
		t.Fail()
	}

	lhs = nil
	rhs = nil
	j = slices.Join(lhs, rhs)
	if j != nil {
		t.Fail()
	}
}

func TestLastIndex(t *testing.T) {
	s := []int{1, 2, 2, 3}

	if slices.LastIndex(s, 2) != 2 {
		t.Fail()
	}

	if slices.LastIndex(s, 0) != -1 {
		t.Fail()
	}

	s = nil
	if slices.LastIndex(s, 1) != -1 {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	s := []int{0, 1, 2}
	f := func(i int) int {
		return i + 1
	}

	mapped := slices.Map(f, s)

	if !slices.Equal(mapped, []int{1, 2, 3}) {
		t.Fail()
	}

	s = nil
	mapped = slices.Map(f, s)
	if mapped != nil {
		t.Fail()
	}
}

func TestRepeat(t *testing.T) {
	s := []int{1, 2, 3}
	r := slices.Repeat(s, 2)

	if !slices.Equal(r, []int{1, 2, 3, 1, 2, 3}) {
		t.Fail()
	}

	r = slices.Repeat(s, 0)
	if r == nil || len(r) != 0 {
		t.Fail()
	}

	s = []int{}
	r = slices.Repeat(s, 2)

	if r == nil || len(s) != 0 {
		t.Fail()
	}

	s = nil
	r = slices.Repeat(s, 2)

	if r != nil {
		t.Fail()
	}
}

func TestReplace(t *testing.T) {
	s := []int{1, 2, 2, 3}
	r := slices.Replace(s, 2, 0, 1)
	if !slices.Equal(r, []int{1, 0, 2, 3}) {
		t.Fail()
	}

	s = []int{1, 2, 2, 3}
	r = slices.Replace(s, 2, 0, 2)
	if !slices.Equal(r, []int{1, 0, 0, 3}) {
		t.Fail()
	}

	s = []int{1, 2, 2, 3}
	r = slices.Replace(s, 2, 0, -1)
	if !slices.Equal(r, []int{1, 0, 0, 3}) {
		t.Fail()
	}

	s = []int{1, 2, 2, 3}
	r = slices.Replace(s, 0, 5, 1)
	if !slices.Equal(r, []int{1, 2, 2, 3}) {
		t.Fail()
	}

	s = []int{1, 2, 2, 3}
	r = slices.Replace(s, 0, 5, -1)
	if !slices.Equal(r, []int{1, 2, 2, 3}) {
		t.Fail()
	}

	s = []int{}
	r = slices.Replace(s, 0, 5, 1)
	if r == nil || len(r) != 0 {
		t.Fail()
	}

	s = []int{}
	r = slices.Replace(s, 0, 5, -1)
	if r == nil || len(r) != 0 {
		t.Fail()
	}

	s = nil
	r = slices.Replace(s, 0, 5, 1)
	if r != nil {
		t.Fail()
	}

	s = nil
	r = slices.Replace(s, 0, 5, -1)
	if r != nil {
		t.Fail()
	}
}

func TestReplaceAll(t *testing.T) {
	s := []int{1, 2, 2, 3}
	r := slices.ReplaceAll(s, 2, 0)
	if !slices.Equal(r, []int{1, 0, 0, 3}) {
		t.Fail()
	}

	s = []int{1, 2, 2, 3}
	r = slices.ReplaceAll(s, 0, 5)
	if !slices.Equal(r, []int{1, 2, 2, 3}) {
		t.Fail()
	}

	s = []int{}
	r = slices.ReplaceAll(s, 0, 5)
	if r == nil || len(r) != 0 {
		t.Fail()
	}

	s = nil
	r = slices.ReplaceAll(s, 0, 5)
	if r != nil {
		t.Fail()
	}
}

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	if !slices.Equal(s1, s1) {
		t.Fail()
	}

	s2 := []int{0, 2, 3}
	if slices.Equal(s1, s2) {
		t.Fail()
	}

	// different len
	s1 = []int{1, 2, 3}
	s2 = []int{1, 2, 3, 4, 5}

	if slices.Equal(s1, s2) {
		t.Fail()
	}

	s1 = nil
	s2 = []int{1, 2, 3}
	if slices.Equal(s1, s2) {
		t.Fail()
	}

	s1 = []int{1, 2, 3}
	s2 = nil
	if slices.Equal(s1, s2) {
		t.Fail()
	}

	s1 = nil
	s2 = nil
	if !slices.Equal(s1, s2) {
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	s := []int{1, 2, 3}
	sum := slices.Sum(s)

	if sum != 1+2+3 {
		t.Fail()
	}

	s = nil
	sum = slices.Sum(s)

	if sum != 0 {
		t.Fail()
	}
}

func TestProd(t *testing.T) {
	s := []int{1, 2, 3}
	prod := slices.Prod(s)

	if prod != 1*2*3 {
		t.Fail()
	}

	s = nil
	prod = slices.Prod(s)

	if prod != 0 {
		t.Fail()
	}
}
