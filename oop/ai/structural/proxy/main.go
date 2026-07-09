package main

import "fmt"

type Server interface {
	HandleRequest(url, method string)
}

type RealServer struct {}

func (s *RealServer) HandleRequest(url, method string) {
	// working
}

type ProxyServer struct {
	realServer *RealServer
}

func (p *ProxyServer) HandleRequest(url, method string) {
	if method == "GET" {
		p.realServer.HandleRequest(url, method)
	} else {
		fmt.Println("Access denied")
	}
}