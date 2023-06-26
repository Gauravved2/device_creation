package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	routes := map[string]string{
		"/api/addDevice": "http://localhost:8001",
	}
	for prefix, target := range routes {
		url, err := url.Parse(target)
		if err != nil {
			log.Fatal(err.Error())
		}
		handler := httputil.NewSingleHostReverseProxy(url)
		fmt.Println(prefix)
		http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})
	}
	http.ListenAndServe(":9000",nil)
}
