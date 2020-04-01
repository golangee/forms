package material

import "net/http"

const css = `
	.titleText {
		font-size: large;
	}
`

func Resources(handler interface{ HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) }) {
	handler.HandleFunc("/style.css", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type","	text/css")
		writer.Write([]byte(css))
	})
}
