package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Установите адрес прокси-сервера
	proxyURL, _ := url.Parse("http://368.188.59.198")

	// Создайте делегата Transport для перенаправления запросов через прокси-сервер
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	// Создайте обработчик, который будет перенаправлять все запросы через прокси-сервер
	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		req.URL.Scheme = proxyURL.Scheme
		req.URL.Host = proxyURL.Host
		req.URL.Path = "/" + req.URL.Path
		fmt.Println("shared query...", req.Header)
	},
		Transport: transport,
	}

	// Запустите сервер на порту 8080 и обрабатывайте все запросы с помощью прокси-сервера
	log.Fatal(http.ListenAndServe(":80", proxy))
}
