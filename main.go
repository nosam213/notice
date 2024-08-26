package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

const version string = "0.1"

var (
	localTime    bool     // display request timestamps in local time (non-default)
	requestCount int  = 0 // start request counting from 0 (default)
)

/*
# [] Display name of time zone for when using local time zone option.
*/

func Site(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	r.ParseMultipartForm(16) // adjustable form in-memory storage, not sure on differences
	// logic for local versus UTC timestamps
	var requestTimestamp string
	if localTime {
		requestTimestamp = time.Now().Format("02-01-2006 15:04:05.000000 Local")
	} else {
		requestTimestamp = time.Now().UTC().Format("02-01-2006 15:04:05.000000 UTC")
	}
	fmt.Printf("\n-- [Request: %d] [%s] --\n", requestCount, requestTimestamp)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Path: %s\n", r.URL.Path)
	// print headers
	for headerName, headerValues := range r.Header {
		for _, headerValue := range headerValues {
			fmt.Printf("%s: %s\n", headerName, headerValue)
		}
	}
	println("~")
	// test printing form data

	for formName, formValue := range r.Form {
		fmt.Printf("%s: %s\n", formName, formValue)
	}
	requestCount++
}

func main() {
	// non global configurations
	var (
		fromOne      bool
		tls          bool
		hideBanner   bool
		port         string = "9001"
		ip           string = "0.0.0.0"
		certLocation string = "./cert.pem"
		keyLocation  string = "./key.pem"
	)
	// flags
	flag.StringVar(&port, "port", port, "the port")
	flag.StringVar(&ip, "ip", ip, "the ip")
	flag.BoolVar(&tls, "tls", false, "TLS (default @ ./cert.pem ./key.pem)")
	flag.BoolVar(&hideBanner, "hide-banner", false, "disables info banner")
	flag.BoolVar(&fromOne, "from-one", false, "starts request counting from 1")
	flag.BoolVar(&localTime, "local-time", false, "enable local time timestamps (default UTC)")
	flag.Parse()

	// request counter (default starts at zero, can be set one)
	if fromOne {
		requestCount = 1
	}
	// routing
	var routing string = fmt.Sprintf("%s:%s", ip, port)
	http.HandleFunc("/", Site)
	// banner (displayed by default)
	if !hideBanner {
		if tls {
			fmt.Printf("notice running at: https://%s [version: %s]\n", routing, version)
		} else {
			fmt.Printf("notice running at: http://%s [version: %s]\n", routing, version)
		}
	}
	// http server (non-tls by default)
	if tls {
		// checking for existance of TLS cert file
		if _, err := os.Stat(certLocation); err == nil {
			// checking for existance of TLS key file
			if _, err := os.Stat(keyLocation); err == nil {
				err := http.ListenAndServeTLS(routing, certLocation, keyLocation, nil)
				if err != nil {
					println("Port already in use or blocked")
					os.Exit(2)
				}
			} else {
				println("TLS cert or key don't exist or aren't accessible")
				os.Exit(2)
			}
		} else {
			println("TLS cert or key don't exist or aren't accessible")
			os.Exit(2)
		}
	} else {
		err := http.ListenAndServe(routing, nil)
		if err != nil {
			println("Port already in use or blocked")
			os.Exit(2)
		}
	}

}
