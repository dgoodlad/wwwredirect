package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sebest/xff"
)

func handler(w http.ResponseWriter, r *http.Request) {
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	www := &url.URL{
		Host:     r.Host,
		Scheme:   protocol,
		Path:     r.URL.Path,
		RawQuery: r.URL.RawQuery,
		Fragment: r.URL.Fragment,
	}
	log.Printf("remote=%s protocol=%s method=%s host=%s path=%s location=%s", r.RemoteAddr, strings.ToUpper(protocol), r.Method, r.Host, r.URL.Path, www.String())

	w.Header().Add("Cache-Control", "public; max-age=31536000")
	http.Redirect(w, r, www.String(), 301)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	iface := os.Getenv("LISTEN_INTERFACE")
	if iface == "" {
		iface = "0.0.0.0"
	}

	log.Printf("Starting wwwredirector service on %s:%s", iface, port)
	http.ListenAndServe(iface+":"+port, xff.Handler(http.HandlerFunc(handler)))
}
