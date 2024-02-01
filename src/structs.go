package main

import (
	"net/http"
	"net/http/httputil"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

type LoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []Server
}
