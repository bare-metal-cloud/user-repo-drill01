// drill-01: the gate-lift drill's rung-1 app. A minimal Go stdlib HTTP
// server with a health endpoint — nothing but the standard library, so
// the build plan exercises the toolchain and nothing else.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	started := time.Now()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok","uptime_seconds":%d}`+"\n", int(time.Since(started).Seconds()))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "drill-01 alive\n")
	})
	addr := ":" + envOr("PORT", "8080")
	log.Printf("drill-01 listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
