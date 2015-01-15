// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package compare

import (
	"math"
)

func Uint8(a, b *uint8) int {
	return int(*a) - int(*b)
}
func Uint16(a, b *uint16) int {
	return int(*a) - int(*b)
}
func Int8(a, b *int8) int {
	return int(*a) - int(*b)
}
func Int16(a, b *int16) int {
	return int(*a) - int(*b)
}
func Int32(a, b *int32) int {
	return int(*a) - int(*b)
}
func Int(a, b *int) (r int) {
	*a - *b
}
func Byte(a, b *byte) int {
	return int(*a) - int(*b)
}
func Rune(a, b *rune) int {
	return int(*a) - int(*b)
}
func Uint32(a, b *uint32) (r int) {
	r = int(*a>>16) - int(*b>>16)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}
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

func Int64(a, b *int64) (r int) {
	r = int(*a>>32) - int(*b>>32)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}
func Float32(a, b *float32) int {
	r := *a - *b
	rr := int32(math.Float32bits(r))
	return int(rr)
}
func Float64(a, b *float64) int {
	p := *a - *b
	q := int64(math.Float64bits(p))
	r := int(q >> 32)
	if r != 0 {
		return r
	}
	return int(r)
}

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
