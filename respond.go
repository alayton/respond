package respond

import (
	"encoding/json"
	"net/http"
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

var (
	htmlContentType = []string{"text/html; charset=utf-8"}
	jsonContentType = []string{"application/json; charset=utf-8"}
	pngContentType  = []string{"image/png"}
	jpgContentType  = []string{"image/jpeg"}
)

// Error is a generic error struct usable for responses
type Error struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// NewError is a helper method for returning an Error object
func NewError(msg string, code int) Error {
	return Error{msg, code}
}

// HTML writes a string to the Writer with a HTML content type
func HTML(w http.ResponseWriter, statusCode int, html string) error {
	w.WriteHeader(statusCode)
	writeContentType(w, htmlContentType)
	_, err := w.Write([]byte(html))
	return err
}

// JSON writes the object to the Writer as a JSON encoded string with a JSON content type
func JSON(w http.ResponseWriter, statusCode int, obj interface{}) error {
	w.WriteHeader(statusCode)
	writeContentType(w, jsonContentType)
	err := json.NewEncoder(w).Encode(obj)
	return err
}

// PNG writes bytes to the Writer with a PNG content type
func PNG(w http.ResponseWriter, statusCode int, buf []byte) error {
	w.WriteHeader(statusCode)
	writeContentType(w, pngContentType)
	_, err := w.Write(buf)
	return err
}

// JPG writes bytes to the Writer with a JPG content type
func JPG(w http.ResponseWriter, statusCode int, buf []byte) error {
	w.WriteHeader(statusCode)
	writeContentType(w, jpgContentType)
	_, err := w.Write(buf)
	return err
}

// NotFound writes an empty response for a 404
func NotFound(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}
