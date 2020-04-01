package wtk

import (
	"fmt"
)

type Color uint32

var Green = RGBA(0, 255, 0, 255)

func RGBA(r, g, b, a byte) Color {
	return Color(uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a))
}

func (c Color) Red() byte {
	return byte(c)
}

func (c Color) Green() byte {
	return byte(c >> 8)
}

func (c Color) Blue() byte {
	return byte(c >> 16)
}

func (c Color) Alpha() byte {
	return byte(c >> 24)
}

func (c Color) String() string {
	return fmt.Sprintf("#%08x", int(c))
}
