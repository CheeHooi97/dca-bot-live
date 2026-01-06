package utils

import (
	"bytes"
	"io"
	"net/http"

	"log"
)

// bodyDumpResponseWriter wraps http.ResponseWriter and copies written data into a buffer
type ResponseWriter struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogRequestWriter(req *http.Request) (reqBody []byte) {
	if req.Body == nil {
		return nil
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Failed to read request body", "error", err)
		return nil
	}

	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	return bodyBytes
}

func LogResponseWriter(original http.ResponseWriter) (*ResponseWriter, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	wrapper := &ResponseWriter{
		ResponseWriter: original,
		Body:           buf,
	}
	return wrapper, buf
}
