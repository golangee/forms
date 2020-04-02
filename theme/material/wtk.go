package material

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

type res struct {
	brotli []byte
	gzip   []byte
}

func Resources(handler interface{ HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) }) {
	files := make(map[string]res)
	files["/material/wtk.css"] = res{mustDecodeBase64(wtkcssbr), mustDecodeBase64(wtkcssgz)}
	files["/material/material-components-web.min.css"] = res{mustDecodeBase64(materialcomponentswebmincssbr), mustDecodeBase64(materialcomponentswebmincssgz)}
	files["/material/material-components-web.min.js"] = res{mustDecodeBase64(materialcomponentswebminjsbr), mustDecodeBase64(materialcomponentswebminjsgz)}
	files["/material/materialicons.woff2"] = res{mustDecodeBase64(materialiconswoffbr), mustDecodeBase64(materialiconswoffgz)}

	handler.HandleFunc("/material/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		switch filepath.Ext(path) {
		case ".css":
			writer.Header().Set("content-type", "text/css")
		case ".js":
			writer.Header().Set("content-type", "application/javascript")
		case ".woff2":
			writer.Header().Set("content-type", "font/woff2")
		default:
			writer.Header().Set("content-type", "application/octet")
		}
		if strings.Contains(request.Header.Get("Accept-Encoding"), "br") {
			writer.Header().Set("Content-Encoding", "br")
			writer.Write(files[path].brotli)
			return
		} else {
			writer.Header().Set("Content-Encoding", "gzip")
			writer.Write(files[path].gzip)
		}

	})

	for fname, data := range files {
		fmt.Printf("providing resource %s providing %d bytes\n", fname, len(data.brotli)+len(data.gzip))
	}

}

func mustDecodeBase64(str string) []byte {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return b
}
