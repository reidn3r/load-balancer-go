package backend

import "net/http/httputil"

type Backend struct {
	Url   string
	Proxy httputil.ReverseProxy
}
