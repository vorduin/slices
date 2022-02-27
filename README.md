# Slices
A utility package for working with slices.

This package resembles the `strings` and `bytes` packages in the Go standard library.

## Installation

Nune requires Go v1.18 as it's entirely based on generics in order to achieve a flexible interface.
Go v1.18 is currently only available in beta version, which can be downloaded [here](https://go.dev/dl/).

After installing Go1.18, simply run this in your terminal...

```zsh
go get github.com/vorduin/nune
```

... and you're good to Go!

## Example Usage

```go
package main

import (
	"github.com/vorduin/slices"
)

func main() {
	s := []int{1, 2, 3}

	clone := slices.Clone(s) // creates a copy and returns it
	slices.Equal(clone, s) // returns true

	slices.Sum(s) // returns the sum of the elems
	slices.Prod(s) // returns the prod of the elems

	slices.Contains(s, 0) // returns whether or not s contains a 0

	incByOne := func(x int) int {
		return x + 1
	}
	s = slices.Map(incByOne, s) // creates a copy of s and maps the function on it
	// s is now [2, 3, 4]

	r := slices.Repeat(s, 2) // repeats s 2 times
	// s is now [2, 3, 4, 2, 3, 4]

	slices.Index(r, 2) // returns 0
	slices.LastIndex(r, 2) // returns 3

	slices.Join(s, r) // returns [2, 3, 4, 2, 3, 4, 2, 3, 4]

	slices.Count(r, 3) // returns the number of 3s in r

	slices.Reverse(s) // returns a copy of s with the elems reversed

	slices.ReplaceAll(s, 4, 0) // returns a copy of s with 4s replaced by 0s
}
```

## License

Slices a BSD-style license, which can be found in the [LICENSE](https://github.com/vorduin/slices/blob/main/LICENSE) file.
