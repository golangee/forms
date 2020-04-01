package wtk

import "strconv"

// assertNotAttached bails out if parent is not nil
func assertNotAttached(v View) {
	if v.parent() != nil {
		panic("invalid state: view is already attached")
	}
}

// assertAttached bails out if parent is nil
func assertAttached(v View) {
	if v.parent() == nil {
		panic("invalid state: view is not attached")
	}
}

func floatToPx(v float64) string {
	return strconv.Itoa(int(v)) + "px"
}
