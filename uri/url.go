package uri

import url2 "net/url"

func ParseURL(s string) (*url2.URL, error) {
	return url2.Parse(s)
}

func Must(u *url2.URL, e error) *url2.URL {
	if e != nil {
		panic(e)
	}
	return u
}
