package backend

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type Backend struct {
	Url   *url.URL
	Proxy *httputil.ReverseProxy
}

func CreateBackend(addr string) *Backend {
	urlPtr, err := url.Parse(addr)

	if err != nil {
		msg := fmt.Sprintf("Url: %s", urlPtr.String())
		fmt.Printf("Erro ao criar novo backend.")
		fmt.Errorf(msg)
	}

	return &Backend{
		Url:   urlPtr,
		Proxy: httputil.NewSingleHostReverseProxy(urlPtr),
	}
}
