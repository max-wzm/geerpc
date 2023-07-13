package geerpc

import (
	"fmt"
	"geerpc/registry"
	"net"
	"net/http"
	"strconv"
	"sync"
)

func startRegistry(wg *sync.WaitGroup, port int) {
	l, _ := net.Listen("tcp", ":"+strconv.Itoa(port))
	registry.HandleHTTP()
	wg.Done()
	_ = http.Serve(l, nil)
}

func StartRegistry(port int) string{
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go startRegistry(wg, port)
	wg.Wait()
	return fmt.Sprintf("http://localhost:%s/_geerpc_/registry", strconv.Itoa(port))
}