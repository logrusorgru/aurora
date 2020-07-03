//
// Copyright (c) 2016-2020 The Aurora Authors. All rights reserved.
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
	"strconv"
	"testing"
)

func TestColor_Nos(t *testing.T) {

	for _, zero := range []bool{
		false, true,
	} {
		for i, val := range []struct {
			Nos   string
			Color Color
		}{
			{"1", BoldFm},
			{"2", FaintFm},
			{"3", ItalicFm},
			{"4", UnderlineFm},
			{"5", SlowBlinkFm},
			{"6", RapidBlinkFm},
			{"7", ReverseFm},
			{"8", ConcealFm},
			{"9", CrossedOutFm},

			{"20", FrakturFm},
			{"21", DoublyUnderlineFm},

			{"51", FramedFm},
			{"52", EncircledFm},
			{"53", OverlinedFm},

			{"30", BlackFg},
			{"31", RedFg},
			{"32", GreenFg},
			{"33", YellowFg},
			{"34", BlueFg},
			{"35", MagentaFg},
			{"36", CyanFg},
			{"37", WhiteFg},

			{"90", BlackFg | BrightFg},
			{"91", RedFg | BrightFg},
			{"92", GreenFg | BrightFg},
			{"93", YellowFg | BrightFg},
			{"94", BlueFg | BrightFg},
			{"95", MagentaFg | BrightFg},
			{"96", CyanFg | BrightFg},
			{"97", WhiteFg | BrightFg},

			{"90", BrightFg},

			{"40", BlackBg},
			{"41", RedBg},
			{"42", GreenBg},
			{"43", YellowBg},
			{"44", BlueBg},
			{"45", MagentaBg},
			{"46", CyanBg},
			{"47", WhiteBg},

			{"100", BlackBg | BrightBg},
			{"101", RedBg | BrightBg},
			{"102", GreenBg | BrightBg},
			{"103", YellowBg | BrightBg},
			{"104", BlueBg | BrightBg},
			{"105", MagentaBg | BrightBg},
			{"106", CyanBg | BrightBg},
			{"107", WhiteBg | BrightBg},

			{"100", BrightBg},

			// bold and faint

			{"1", BoldFm | FaintFm},

			// slow blink and rapid blink

			{"5", SlowBlinkFm | RapidBlinkFm},

			// index

			{"38;5;100", (100 << shiftFg) | flagFg},
			{"48;5;100", (100 << shiftBg) | flagBg},

			// longest combination

			{"1;3;4;5;7;8;9;20;21;51;52;53;38;5;123;48;5;231",
				BoldFm | FaintFm |
					ItalicFm | UnderlineFm |
					SlowBlinkFm | RapidBlinkFm |
					ReverseFm | ConcealFm |
					CrossedOutFm | FrakturFm | DoublyUnderlineFm |
					FramedFm | EncircledFm | OverlinedFm |
					Color(123)<<shiftFg | flagFg |
					Color(231)<<shiftBg | flagBg},
		} {
			var (
				nos  = val.Color.Nos(zero)
				want = val.Nos
			)
			if zero {
				if want != "" {
					want = "0;" + want
				} else {
					want = "0"
				}
			}
			if nos != want {
				t.Errorf("%t %d: wrong nos string %q, want %q",
					zero, i, nos, want)
			}
		}
	}
}

func TestColor_IsValid(t *testing.T) {
	if Color(0).IsValid() == false {
		t.Error("invalid")
	}
}

func Test_itoa(t *testing.T) {
	for i := 0; i < 256; i++ {
		if a := itoa(byte(i)); a != strconv.Itoa(i) {
			t.Errorf("wrong %q, want %d", a, i)
		}
	}
}
