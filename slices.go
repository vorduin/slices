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
	cp := make([]T, len(s))
	copy(cp, s)

	return cp
}

// Equal returns whether or not two slices are equal.
func Equal[T constraints.Ordered](a, b []T) bool {
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
	var prod T = 1
	for i := 0; i < len(s); i++ {
		prod *= s[i]
	}
	return prod
}
