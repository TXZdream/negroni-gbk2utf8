package gbk2utf8

import (
	"github.com/urfave/negroni"
	"io/ioutil"
	"strings"
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	ContentType = "Content-Type"
)

type gtuResponseWriter struct {
	w *transform.Writer
	negroni.ResponseWriter
	wroteHeader bool
}

type handler struct {}

func Transformer() *handler {
	h := &handler{}
	return h
}

// Set Content-Type in response header
// Skip if there exists another encoding
func (gtu *gtuResponseWriter) WriteHeader(code int) {
	headers := gtu.ResponseWriter.Header()
	if len(headers.Get(ContentType)) == 0 || strings.Contains(headers.Get(ContentType), "UTF-8") {
		headers.Set(ContentType, strings.Replace(headers.Get(ContentType), "UTF-8", "gbk", -1))
	} else {
		gtu.w = nil
	}
	gtu.ResponseWriter.WriteHeader(code)
	gtu.wroteHeader = true
}

// Write encoding byte
// Set http response header
func (gtu *gtuResponseWriter) Write(b []byte) (int, error) {
	if gtu.w == nil {
		return gtu.ResponseWriter.Write(b)
	}
	if len(gtu.Header().Get(ContentType)) == 0 {
		gtu.Header().Set(ContentType, http.DetectContentType(b))
	}
	return gtu.w.Write(b)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Set Content-Type to UTF-8 if not exists
	if !strings.Contains(r.Header.Get(ContentType), "gbk") {
		r.Header.Set(ContentType, "UTF-8")
	}
	
	// Read from http body
	// Transform gbk encoding to utf8 encoding
	// Reconstruct http body and replace the raw one
	// Replace header
	// Set writer to encoding the content to gbk
	reader := transform.NewReader(r.Body, simplifiedchinese.GBK.NewDecoder())
	r.Body = ioutil.NopCloser(reader)
	r.Header.Set(ContentType, strings.Replace(r.Header.Get(ContentType), "gbk", "UTF-8", -1))

	// Wrap writer to encoding content
	nw := negroni.NewResponseWriter(w)
	writer := transform.NewWriter(nw, simplifiedchinese.GBK.NewEncoder())
	gw := gtuResponseWriter{writer, nw, false}
	next(&gw, r)
}