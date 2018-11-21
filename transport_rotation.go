package main

import (
	"log"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

func updateTransport(t selector.ProxyType, proxy *goproxy.ProxyHttpServer, maxTries int) error {
	log.Println("Rotation started...")
	transport, err := selector.GetNextProxyTransport(t, maxTries)
	if err != nil {
		return err
	}

	proxy.Tr = transport
	log.Println("Rotation finished...")
	return nil
}

func rotateTransport(t selector.ProxyType, proxy *goproxy.ProxyHttpServer, timeout time.Duration, maxTries int) {
	for {
		if timeout == 0 {
			break
		}
		time.Sleep(timeout)
		err := updateTransport(t, proxy, maxTries)
		if err != nil {
			log.Printf("Transport rotation error: %s", err)
		}
	}
}
