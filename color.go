package wtk

import (
	"fmt"
)

type Color uint32

func RGBA(r, g, b, a byte) Color {
	return Color(uint32(r) | uint32(g)<<8 | uint32(b)<<16 | uint32(a)<<24)
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
