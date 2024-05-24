// adapted from https://github.com/justinas/alice
package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareProxy struct {
	mids []Middleware
}

func New(mids ...Middleware) MiddlewareProxy {
	midsArr := make([]Middleware, len(mids)-1)
	midsArr = append(midsArr, mids...)
	return MiddlewareProxy{mids: midsArr}
}

func (mp MiddlewareProxy) Then(h http.Handler) http.Handler {
	mid := h
	if mid == nil {
		mid = http.DefaultServeMux
	}

	for i := range mp.mids {
		// fmt.Println(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
		mid = mp.mids[len(mp.mids)-i-1](h)
	}
	return mid
}

func (mp MiddlewareProxy) ThenFunc(fn http.HandlerFunc) http.Handler {
	// Required due to: https://stackoverflow.com/questions/33426977/how-to-golang-check-a-variable-is-nil
	if fn == nil {
		return mp.Then(nil)
	}
	return mp.Then(fn)
}

func (mp MiddlewareProxy) Extend(mids ...Middleware) MiddlewareProxy {
	midsArr := make([]Middleware, len(mids)+len(mp.mids))
	midsArr = append(midsArr, mp.mids...)
	midsArr = append(midsArr, mids...)
	return MiddlewareProxy{mids: midsArr}
}
