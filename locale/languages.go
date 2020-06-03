package locale

import "syscall/js"

var overrideLanguages []string

// SetLanguages overrides the system languages. Set to nil/empty to reset and let the system preference take precedence.
func SetLanguages(preferredLanguages ...string) {
	overrideLanguages = preferredLanguages
}

// Languages returns at least a single element array with "und" (undefined) value, but actually represents the
// users preferred system languages.
func Languages() []string {
	if len(overrideLanguages) > 0 {
		return overrideLanguages
	}

	arr := js.Global().Get("navigator").Get("languages")
	res := make([]string, arr.Length())
	for i := 0; i < arr.Length(); i++ {
		res[i] = arr.Index(i).String()
	}

	if len(res) == 0 {
		res = append(res, "und")
	}

	return res
}

// Language returns the preferred user language.
func Language() string {
	return Languages()[0]
}
