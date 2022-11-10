package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

var (
	baseURL = "http://localhost:808"
)
type LoadBalancer struct {
	RevProxy httputil.ReverseProxy
}

type EndPoints struct{
	List []*url.URL
}
