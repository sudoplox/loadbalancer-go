package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// NewLoadBalancer returns a new LoadBalancer
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers,
	}
}

// SimpleServer is a basic implementation of the Server interface
func newSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) IsAlive() bool {
	// TODO: Implement health check
	return true
}

// Serve proxies the request to the server
func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(rw, req)
}

// GetNextAvailableServer returns the next available server in the list of servers
func (lb *LoadBalancer) GetNextAvailableServer() Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}

// ServeProxy proxies the request to the next available server
func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Printf("Proxying request to: %v\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func main() {
	servers := []Server{
		newSimpleServer("http://www.facebook.com"),
		newSimpleServer("http://www.google.com"),
		newSimpleServer("http://www.linkedin.com"),
	}

	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Load Balancer listening on port: %s\n", lb.Port)
	err := http.ListenAndServe(":"+lb.Port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
