// Copyright © The Slices Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import "golang.org/x/exp/constraints"

// WithLen makes a new slice with the given length and returns it.
func WithLen[T any](l int) []T {
	return make([]T, l)
}

// WithCap makes a new slice with the given capacity and returns it.
func WithCap[T any](c int) []T {
	return make([]T, 0, c)
}

// Clone makes a new slice and copies the given slice's elements into it.
func Clone[T any](s []T) []T {
	if s == nil {
		return nil
	}

	cp := make([]T, len(s))
	copy(cp, s)

	return cp
}

// Reverse returns a copy of the slice s with order of the elements reversed.
func Reverse[T any](s []T) []T {
	cp := Clone(s)

	for i, j := 0, len(cp)-1; i < j; i, j = i+1, j-1 {
		cp[i], cp[j] = cp[j], cp[i]
	}

	return cp
}

// Last returns the last element in a slice.
// Panics if s is empty.
func Last[T any](s []T) T {
	if len(s) == 0 {
		panic("slices: empty slice")
	}

	return s[len(s)-1]
}

// Contains returns whether or not a slice contains a certain element.
func Contains[T constraints.Ordered](s []T, el T) bool {
	for _, x := range s {
		if x == el {
			return true
		}
	}

	return false
}

// Count counts the number of instances of el in s
func Count[T constraints.Ordered](s []T, el T) int {
	var c int
	for _, x := range s {
		if x == el {
			c += 1
		}
	}

	return c
}

// Index returns the index of the first instance of el in s,
// or -1 if el is not present in s.
func Index[T constraints.Ordered](s []T, el T) int {
	for i, x := range s {
		if x == el {
			return i
		}
	}

	return -1
}

// Join joins two slices together.
// If an argument is nil, it returns a clone
// of the other argument.
// If both are nil, it returns nil.
func Join[T any](lhs, rhs []T) []T {
	if lhs == nil && rhs == nil {
		return nil
	} else if lhs == nil {
		return Clone(rhs)
	} else if rhs == nil {
		return Clone(lhs)
	}

	s := make([]T, len(lhs) + len(rhs))
	copy(s, lhs)
	copy(s[len(lhs):], rhs)

	return s
}

// LastIndex returns the index of the last instance of el in s,
// or -1 if el is not present in s. 
func LastIndex[T constraints.Ordered](s []T, el T) int {
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == el {
			return i
		}
	}

	return -1
}

// Map returns a copy of the slice s with all its elements modified
// according to the mapping function.
func Map[T any](mapping func(T) T, s []T) []T {
	cp := Clone(s)

	for i := 0; i < len(s); i++ {
		cp[i] = mapping(s[i])
	}

	return cp
}

// Repeat returns a new slice consisting of count copies of the slice s.
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func Repeat[T any](s []T, count int) []T {
	if s == nil {
		return nil
	}

	if count == 0 {
		return []T{}
	}

	if count < 0 {
		panic("slices: negative Repeat count")
	} else if len(s)*count/count != len(s) {
		panic("slices: Repeat count causes overflow")
	}

	res := make([]T, len(s) * count)
	for i := 0; i < count; i++ {
		copy(res[i*len(s):(i+1)*len(s)], s)
	}

	return res
}

// Replace returns a copy of the slice s with the first n instances
// of old replaced by new. If n < 0, there is no limit on
// the number of replacements.
func Replace[T constraints.Ordered](s []T, old, new T, n int) []T {
	if s == nil {
		return nil
	}

	cp := make([]T, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == old && n != 0 {
			cp[i] = new
			n--
		} else {
			cp[i] = s[i]
		}
	}

	return cp
}

// ReplaceAll returns a copy of the slice s with all instances of
// old replaced by new.
func ReplaceAll[T constraints.Ordered](s []T, old, new T) []T {
	if s == nil {
		return nil
	}

	cp := make([]T, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == old {
			cp[i] = new
		} else {
			cp[i] = s[i]
		}
	}

	return cp
}

// Equal returns whether or not two slices are equal.
func Equal[T constraints.Ordered](a, b []T) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Sum returns the sum of the elements of the slice.
func Sum[T constraints.Integer | constraints.Float](s []T) T {
	var sum T
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

// Prod returns the product of the elements of the slice.
func Prod[T constraints.Integer | constraints.Float](s []T) T {
	if len(s) == 0 {
		return 0
	}

	var prod T = 1
	for i := 0; i < len(s); i++ {
		prod *= s[i]
	}
	return prod
}
