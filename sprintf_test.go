//
// Copyright (c) 2016-2022 The Aurora Authors. All rights reserved.
// This program is free software. It comes without any warranty,
// to the extent permitted by applicable law. You can redistribute
// it and/or modify it under the terms of the Unlicense. See LICENSE
// file for more details or see below.
//

//
// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org/>
//

package aurora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type noString string

func Test_Sprintf(t *testing.T) {
	var want, got string

	want = "delta: +3"
	got = Sprintf(noString("delta: +%d"), 3)
	assert.Equal(t, want, got)

	want = "\033[31mdeltas: +3, 5, 9 points\033[0m"
	got = Sprintf(Red("deltas: +%d, %d, %d points"), 3, 5, 9)
	assert.Equal(t, want, got)

	// %s
	want = `[31mquoted: [0;34m    "blue"[0;31m` +
		`, [0;32m     green[0;31m[0m`
	got = Sprintf(Red("quoted: % 10q, % 10s"), Blue("blue"), Green("green"))
	assert.Equal(t, want, got)

	// on string
	want = `quoted: [34mblue[0m, [32mgreen[0m`
	got = Sprintf("quoted: %s, %s", Blue("blue"), Green("green"))
	assert.Equal(t, want, got)

	// no tail
	want = `quoted: [34mblue[0m, [32mgreen[0m`
	got = Sprintf(Clear("quoted: %s, %s"), Blue("blue"), Green("green"))
	assert.Equal(t, want, got)

	// precision
	want = `[31mvalue: [0;34m2.78[0;31m[0m`
	got = Sprintf(Red("value: %1.2f"), Blue(2.7834))
	assert.Equal(t, want, got)

	// wide verb
	want = `[31m[0;34m%!世(float64=+2.78)[0;31m[0m`
	got = Sprintf(Red("%+1.3世"), Blue(2.7834))
	assert.Equal(t, want, got)

	// decolor
	var au = New(WithColors(false), WithHyperlinks(false))
	want = `+2.783`
	got = au.Sprintf(Red("%+1.3f"), Blue(2.7834))
	assert.Equal(t, want, got)
}
