// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

// Package compare provides a compare functions to compare two objects.
//
// Returns 0 for equality.
//
// Returns - (negative) for a < b.
//
// Returns + (positive) for a > b.
package compare

import (
	"math"
	"math/cmplx"
)

// compare.Uint8 compares two numbers.
func Uint8(a, b *uint8) int {
	return int(*a) - int(*b)
}

// compare.Uint16 compares two numbers.
func Uint16(a, b *uint16) int {
	return int(*a) - int(*b)
}

// compare.Int8 compares two numbers.
func Int8(a, b *int8) int {
	return int(*a) - int(*b)
}

// compare.Int16 compares two numbers.
func Int16(a, b *int16) int {
	return int(*a) - int(*b)
}

// compare.Int32 compares two numbers.
func Int32(a, b *int32) int {
	return int(*a) - int(*b)
}

// compare.Int compares two numbers.
func Int(a, b *int) (r int) {
	return *a - *b
}

// compare.Byte compares two numbers.
func Byte(a, b *byte) int {
	return int(*a) - int(*b)
}

// compare.Rune compares two numbers.
func Rune(a, b *rune) int {
	return int(*a) - int(*b)
}

// compare.Uint32 compares two numbers.
func Uint32(a, b *uint32) (r int) {
	r = int(*a>>16) - int(*b>>16)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}

// compare.Uint64 compares two numbers.
func Uint64(a, b *uint64) (r int) {
	r = int(*a>>32) - int(*b>>32)
	if r != 0 {
		return r
	}
	r = int(*a>>16) - int(*b>>16)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}

// compare.Int64 compares two numbers.
func Int64(a, b *int64) (r int) {
	r = int(*a>>32) - int(*b>>32)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}

// compare.Float32 compares two numbers.
func Float32(a, b *float32) int {
	r := *a - *b
	rr := int32(math.Float32bits(r))
	return int(rr)
}

// compare.Float64 compares two numbers.
func Float64(a, b *float64) int {
	p := *a - *b
	q := int64(math.Float64bits(p))
	r := int(q >> 32)
	if r != 0 {
		return r
	}
	return int(r)
}

// compare.Uint compares two numbers.
func Uint(a, b *uint) (r int) {
	r = int(*a>>32) - int(*b>>32)
	if r != 0 {
		return r
	}
	r = int(*a>>16) - int(*b>>16)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}

// compare.Abs compares two complex numbers by absolute value.
func Abs(a, b *complex128) int {
	fa := cmplx.Abs(*a)
	fb := cmplx.Abs(*b)
	p := fa - fb
	q := int64(math.Float64bits(p))
	r := int(q >> 32)
	if r != 0 {
		return r
	}
	return int(r)
}

// compare.Real compares two complex numbers by real part.
func Real(a, b *complex128) int {
	fa := real(*a)
	fb := real(*b)
	p := fa - fb
	q := int64(math.Float64bits(p))
	r := int(q >> 32)
	if r != 0 {
		return r
	}
	return int(r)
}

// compare.Imag compares two complex numbers by imaginary part.
func Imag(a, b *complex128) int {
	fa := imag(*a)
	fb := imag(*b)
	p := fa - fb
	q := int64(math.Float64bits(p))
	r := int(q >> 32)
	if r != 0 {
		return r
	}
	return int(r)
}
