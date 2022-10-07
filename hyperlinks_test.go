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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHyperlinkParam_stringLen(t *testing.T) {
	var param HyperlinkParam
	assert.Equal(t, 1, param.stringLen())
	param.Key, param.Value = "some", "thing"
	assert.Equal(t, 1+len("some")+len("thing"), param.stringLen())
}

func TestHyperlinkParam_String(t *testing.T) {
	var param HyperlinkParam
	assert.Equal(t, "=", param.String())
	param.Key, param.Value = "some", "thing"
	assert.Equal(t, "some=thing", param.String())
}

func TestIsValidHyperlinkTarget(t *testing.T) {
	assert.True(t, IsValidHyperlinkTarget("http://example.com/path?query=true"))
	assert.True(t, IsValidHyperlinkTarget("mailto:user@example.com"))
	assert.False(t, IsValidHyperlinkTarget("http://пример.тест/путь?запрос=да"))
	assert.False(t, IsValidHyperlinkTarget("mailto:пользователь@пример.тест"))
}

func TestIsValidHyperlinkParam(t *testing.T) {
	assert.True(t, IsValidHyperlinkParam("id"))
	assert.True(t, IsValidHyperlinkParam("a_param-key"))
	assert.False(t, IsValidHyperlinkParam("value:true"))
	assert.False(t, IsValidHyperlinkParam("value=true"))
	assert.False(t, IsValidHyperlinkParam("v1;v2"))
}

func TestHyperlinkID(t *testing.T) {
	assert.Equal(t, HyperlinkParam{
		Key:   HyperlinkIDKey,
		Value: "value",
	}, HyperlinkID("value"))
}

// TODO
//
// func (h *hyperlink) isExists() (ok bool)
// func (h *hyperlink) stringParamsLen() (ln int)
// func (h *hyperlink) headLen() int
// func (h *hyperlink) headBytes() (t []byte)
// func (h *hyperlink) head() string
// func (h *hyperlink) tailLen() int
// func (h *hyperlink) tailBytes() []byte
// func (h *hyperlink) tail() string
// func (h *hyperlink) writeHead(w io.Writer)
// func (h *hyperlink) writeTail(w io.Writer)

func Test_unhex(t *testing.T) {
	assert.Zero(t, unhex(0))
}

func TestHyperlinkEscape(t *testing.T) {
	var val = "http://example.com/path?query=true"
	assert.Equal(t, val, HyperlinkEscape(val))
	val = "mailto:user@example.com"
	assert.Equal(t, val, HyperlinkEscape(val))
	val = "http://пример.тест/путь?запрос=да"
	assert.True(t, IsValidHyperlinkTarget(HyperlinkEscape(val)))
	val = "mailto:пользователь@пример.тест"
	assert.True(t, IsValidHyperlinkTarget(HyperlinkEscape(val)))
}

func TestHyperlinkUnescape(t *testing.T) {
	for _, val := range []string{
		"http://example.com/path?query=true",
		"mailto:user@example.com",
		"http://пример.тест/путь?запрос=да",
		"mailto:пользователь@пример.тест",
	} {
		var got, err = HyperlinkUnescape(HyperlinkEscape(val))
		require.NoError(t, err)
		assert.Equal(t, val, got)
	}
	var _, err = HyperlinkUnescape("%%%%")
	assert.Error(t, err)
	var (
		val     = "значение"
		escaped = HyperlinkEscape(val)
		back    string
	)
	escaped = strings.ToLower(escaped)
	back, err = HyperlinkUnescape(escaped)
	assert.NoError(t, err)
	assert.Equal(t, val, back)
}
