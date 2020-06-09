// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package forms

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

type Color uint32

var Green = RGBA(0, 255, 0, 255)

// ParseColor currently can only evaluate codes like #6200ee??
func ParseColor(text string) Color {
	text = strings.TrimSpace(text)
	log.Println("!!!should parse ", text)
	if strings.HasPrefix(text, "#") {
		text = text[1:]
		for len(text) < 8 {
			text = text + "F"
		}
		log.Println("should parse ", text)

		channels, _ := hex.DecodeString(text)
		if len(channels) == 4 {
			return RGBA(channels[0], channels[1], channels[2], channels[3])
		}
	}
	return 0
}

func RGBA(r, g, b, a byte) Color {
	return Color(uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a))
}

func RGB(r, g, b byte) Color {
	return RGBA(r, g, b, 255)
}

func (c Color) Alpha() byte {
	return byte(c)
}

func (c Color) Blue() byte {
	return byte(c >> 8)
}

func (c Color) Green() byte {
	return byte(c >> 16)
}

func (c Color) Red() byte {
	return byte(c >> 24)
}

func (c Color) String() string {
	return fmt.Sprintf("#%08x", int(c))
}

func (c Color) SetAlpha(a byte) Color {
	return RGBA(c.Red(), c.Green(), c.Blue(), a)
}

const (
	Purple       Color = 0x6200EEFF
	Teal         Color = 0x03DAC5FF
	Error        Color = 0xB00020FF
	Red50        Color = 0xF44336FF
	Pink50       Color = 0xE91E63FF
	Purple50     Color = 0x9C27B0FF
	DeepPurple50 Color = 0x673AB7FF
	Indigo50     Color = 0x3F51B5FF
	Blue50       Color = 0x2196F3FF
	LightBlue    Color = 0x03A9F4FF
	Cyan50       Color = 0x00BCD4FF
	Teal50       Color = 0x009688FF
	Green50      Color = 0x4CAF50FF
	LightGreen50 Color = 0x8BC34AFF
	Lime50       Color = 0xCDDC39FF
	Yellow50     Color = 0xFFEB3BFF
	Amber50      Color = 0xFFC107FF
	Orange50     Color = 0xFF9800FF
	DeepOrange50 Color = 0xFF5722FF
	Brown50      Color = 0x795548
	Gray50       Color = 0x9e9e9e
	BlueGray50   Color = 0x607D8BFF
)

var Colors = map[string]Color{
	"Purple":       Purple,
	"Teal":         Teal,
	"Error":        Error,
	"Red50":        Red50,
	"Pink50":       Pink50,
	"Purple50":     Purple50,
	"DeepPurple50": DeepPurple50,
	"Indigo50":     Indigo50,
	"Blue50":       Blue50,
	"LightBlue":    LightBlue,
	"Cyan50":       Cyan50,
	"Teal50":       Teal50,
	"Green50":      Green50,
	"LightGreen50": LightGreen50,
	"Lime50":       Lime50,
	"Yellow50":     Yellow50,
	"Amber50":      Amber50,
	"Orange50":     Orange50,
	"DeepOrange50": DeepOrange50,
	"Brown50":      Brown50,
	"Gray50":       Gray50,
	"BlueGray50":   BlueGray50,
}
