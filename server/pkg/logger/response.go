package logger

import (
	"bytes"
	"net/http"
)

type response struct {
	body       bytes.Buffer
	statusCode int
}

func (r *response) Header() http.Header {
	return http.Header{}
}

func (r *response) Write(b []byte) (int, error) {
	r.body.Write(b)
	return 0, nil
}

func (r *response) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}
