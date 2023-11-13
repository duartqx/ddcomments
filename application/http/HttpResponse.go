package http

import "net/http"

type HttpResponse struct {
	Status int
	Body   interface{}
	Cookie *http.Cookie
}
