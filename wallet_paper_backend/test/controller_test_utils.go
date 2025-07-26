package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

// PerformRequest 发起一个测试 HTTP 请求
func PerformRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
    req := httptest.NewRequest(method, path, bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    return w
}
