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

//
// Foreground colors
//

// Black convertes argument into formated/colorized value
func Black(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | BlackFg
		return val
	}
	return Value{arg, BlackFg, 0}
}

// Red convertes argument into formated/colorized value
func Red(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | RedFg
		return val
	}
	return Value{arg, RedFg, 0}
}

// Green convertes argument into formated/colorized value
func Green(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | GreenFg
		return val
	}
	return Value{arg, GreenFg, 0}
}

// Brown convertes argument into formated/colorized value
func Brown(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | BrownFg
		return val
	}
	return Value{arg, BrownFg, 0}
}

// Blue convertes argument into formated/colorized value
func Blue(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | BlueFg
		return val
	}
	return Value{arg, BlueFg, 0}
}

// Magenta convertes argument into formated/colorized value
func Magenta(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | MagentaFg
		return val
	}
	return Value{arg, MagentaFg, 0}
}

// Cyan convertes argument into formated/colorized value
func Cyan(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | CyanFg
		return val
	}
	return Value{arg, CyanFg, 0}
}

// Gray convertes argument into formated/colorized value
func Gray(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskFg)) | GrayFg
		return val
	}
	return Value{arg, GrayFg, 0}
}

//
// Background colors
//

// BgBlack convertes argument into formated/colorized value
func BgBlack(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | BlackBg
		return val
	}
	return Value{arg, BlackBg, 0}
}

// BgRed convertes argument into formated/colorized value
func BgRed(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | RedBg
		return val
	}
	return Value{arg, RedBg, 0}
}

// BgGreen convertes argument into formated/colorized value
func BgGreen(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | GreenBg
		return val
	}
	return Value{arg, GreenBg, 0}
}

// BgBrown convertes argument into formated/colorized value
func BgBrown(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | BrownBg
		return val
	}
	return Value{arg, BrownBg, 0}
}

// BgBlue convertes argument into formated/colorized value
func BgBlue(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | BlueBg
		return val
	}
	return Value{arg, BlueBg, 0}
}

// BgMagenta convertes argument into formated/colorized value
func BgMagenta(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | MagentaBg
		return val
	}
	return Value{arg, MagentaBg, 0}
}

// BgCyan convertes argument into formated/colorized value
func BgCyan(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | CyanBg
		return val
	}
	return Value{arg, CyanBg, 0}
}

// BgGray convertes argument into formated/colorized value
func BgGray(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color = (val.color & (^maskBg)) | GrayBg
		return val
	}
	return Value{arg, GrayBg, 0}
}

// Bold convertes argument into formated/colorized value
func Bold(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color |= BoldFm
		return val
	}
	return Value{arg, BoldFm, 0}
}

// Inverse convertes argument into formated/colorized value
func Inverse(arg interface{}) Value {
	if val, ok := arg.(Value); ok {
		val.color |= InverseFm
		return val
	}
	return Value{arg, InverseFm, 0}
}
