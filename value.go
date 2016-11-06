//
// Copyright (c) 2016 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE.md file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package aurora

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// A Value represents any printable value
// with it's color
type Value struct {
	value interface{}
	color Color
	tail  Color // format after the value
}

// String returns string with colors. If there are any color
// or format the string will be terminated with \033[0m
func (v Value) String() string {
	if v.color != 0 && v.color.IsValid() {
		if v.tail != 0 && v.tail.IsValid() {
			return esc + v.color.Nos() + "m" + fmt.Sprint(v.value) + clear +
				esc + v.tail.Nos() + "m"
		}
		return esc + v.color.Nos() + "m" + fmt.Sprint(v.value) + clear
	}
	return fmt.Sprint(v.value)
}

// Color returns value's color
func (v Value) Color() Color { return v.color }

// Bleach returns copy of orignal value without colors
func (v Value) Bleach() Value {
	v.color, v.tail = 0, 0
	return v
}

// Value returns value's value (welcome to the tautology club)
func (v Value) Value() interface{} { return v.value }

// Format implements fmt.Formater interface
func (v Value) Format(s fmt.State, verb rune) {
	// \033[1;7;31;45m   - 12 (+tail)
	// %-+# 025.36s      - 12
	// \033[0m           - 4
	var format = getFormat()
	var colors bool
	if v.color != 0 && v.color.IsValid() {
		colors = true
		format = append(format, esc...)
		format = append(format, v.color.Nos()...)
		format = append(format, 'm')
	}
	format = append(format, '%')
	var f byte
	for i := 0; i < len(availFlags); i++ {
		if f = availFlags[i]; s.Flag(int(f)) {
			format = append(format, f)
		}
	}
	var width, prec int
	var ok bool
	if width, ok = s.Width(); ok {
		format = strconv.AppendInt(format, int64(width), 10)
	}
	if prec, ok = s.Precision(); ok {
		format = append(format, '.')
		format = strconv.AppendInt(format, int64(prec), 10)
	}
	if verb > utf8.RuneSelf {
		format = append(format, string(verb)...)
	} else {
		format = append(format, byte(verb))
	}
	if colors {
		format = append(format, clear...)
		if v.tail != 0 && v.tail.IsValid() { // next format
			format = append(format, esc...)
			format = append(format, v.tail.Nos()...)
			format = append(format, 'm')
		}
	}
	fmt.Fprintf(s, string(format), v.value)
	putFormat(format)
}
