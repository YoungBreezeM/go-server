package utils

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func CreateProxyHTTPClient(proxyURL string) *http.Client {
	// 解析代理URL
	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		log.Fatal("Failed to parse proxy URL:", err)
	}

	// 创建Dialer函数
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	// 创建Transport
	transport := &http.Transport{
		Dial:                dialer.Dial,
		Proxy:               http.ProxyURL(proxyURLParsed),
		TLSHandshakeTimeout: 10 * time.Second,
	}

	// 创建HTTP客户端
	httpClient := &http.Client{
		Transport: transport,
	}

	return httpClient
}
