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
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	assert.Equal(t, Config{
		Colors:     true,
		Hyperlinks: true,
	}, NewConfig())
}

func TestConfig_AddFlags(t *testing.T) {
	// set to true
	var fset = flag.NewFlagSet("x", flag.ContinueOnError)
	var conf = NewConfig()
	conf.AddFlags(fset, "testing.")
	var err = fset.Parse([]string{
		"-testing.colors",
		"-testing.hyperlinks",
	})
	require.NoError(t, err)
	assert.True(t, conf.Colors)
	assert.True(t, conf.Hyperlinks)
	// set to true & false
	fset = flag.NewFlagSet("x", flag.ContinueOnError)
	conf.AddFlags(fset, "testing.")
	err = fset.Parse([]string{
		"-testing.colors=t",
		"-testing.hyperlinks=f",
	})
	require.NoError(t, err)
	assert.True(t, conf.Colors)
	assert.False(t, conf.Hyperlinks)
	// set to false & true
	fset = flag.NewFlagSet("x", flag.ContinueOnError)
	conf.AddFlags(fset, "testing.")
	err = fset.Parse([]string{
		"-testing.colors=f",
		"-testing.hyperlinks=t",
	})
	require.NoError(t, err)
	assert.False(t, conf.Colors)
	assert.True(t, conf.Hyperlinks)
	// set to false
	fset = flag.NewFlagSet("x", flag.ContinueOnError)
	conf.AddFlags(fset, "testing.")
	err = fset.Parse([]string{
		"-testing.colors=f",
		"-testing.hyperlinks=f",
	})
	require.NoError(t, err)
	assert.False(t, conf.Colors)
	assert.False(t, conf.Hyperlinks)
}

func TestConfig_Apply(t *testing.T) {
	var conf = NewConfig()
	conf.Apply(WithColors(false), WithHyperlinks(false))
	assert.Equal(t, Config{
		Colors:     false,
		Hyperlinks: false,
	}, conf)
}

func TestConfig_Options(t *testing.T) {
	var (
		c1 = NewConfig()
		c2 Config
	)
	c2.Apply(c1.Options()...)
	assert.Equal(t, Config{
		Colors:     true,
		Hyperlinks: true,
	}, c2)
}

func TestConfig_colorConfig(t *testing.T) {
	var conf = NewConfig()
	assert.Equal(t, colorPin|hyperlinksPin, conf.colorConfig())
	conf.Colors = false
	assert.Equal(t, hyperlinksPin, conf.colorConfig())
	conf.Hyperlinks = false
	assert.Equal(t, colorConfig(0), conf.colorConfig())
}

func TestWithColors(t *testing.T) {
	var conf Config
	// turn to true
	conf.Apply(WithColors(true))
	assert.Equal(t, Config{
		Colors:     true,
		Hyperlinks: false,
	}, conf)
	// turn to false
	conf.Apply(WithColors(false))
	assert.Equal(t, Config{
		Colors:     false,
		Hyperlinks: false,
	}, conf)
}

func TestWithHyperlinks(t *testing.T) {
	var conf Config
	// turn to true
	conf.Apply(WithHyperlinks(true))
	assert.Equal(t, Config{
		Colors:     false,
		Hyperlinks: true,
	}, conf)
	// turn to false
	conf.Apply(WithHyperlinks(false))
	assert.Equal(t, Config{
		Colors:     false,
		Hyperlinks: false,
	}, conf)
}
