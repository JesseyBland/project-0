package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	// Reverse Proxy points from the load balancer to the reverse proxy
	fmt.Println("Reverse Proxy Initializing...")
	proxy := httputil.NewSingleHostReverseProxy(loadBal())
	http.ListenAndServe(":8081", proxy)

}

// Load Balancer evenly distrubted load between 3 servers.
func loadBal() *url.URL {
	var a, b, c int

	if a == b && a == c {
		a++
		fmt.Println("On Server :8888   a=", a)
		ticURL, err := url.Parse("http://localhost:8888")
		if err != nil {
			log.Fatal(err)
		}
		return ticURL

	} else if a > b && b == c {
		b++
		fmt.Println("On Server :9999   b=", b)
		ticURL, err1 := url.Parse("http://localhost:9999")
		if err1 != nil {
			log.Fatal(err1)
		}
		return ticURL

	} else if a > c && b > c {
		c++
		fmt.Println("On Server :7777   c=", c)
		ticURL, err2 := url.Parse("http://localhost:7777")
		if err2 != nil {
			log.Fatal(err2)
		}
		return ticURL
	}

	return nil

}
