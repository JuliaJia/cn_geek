package main

import (
	"log"
	"math/rand"
	"metrics"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", httpRootAccess)
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(rw http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(3000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	rw.Write([]byte("200"))
}

func httpRootAccess(rw http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			rw.Header().Set(k, v[0])
		}
	}
	rw.Write([]byte("Hello,Geeker!"))
}
