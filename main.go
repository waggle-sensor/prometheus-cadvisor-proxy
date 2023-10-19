package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var timeout = flag.Duration("timeout", 30*time.Second, "timeout for scraping node metrics")

func handler(w http.ResponseWriter, r *http.Request) {
	nodeID, _, _ := strings.Cut(r.URL.Path, "/")
	nodeID = strings.ToLower(nodeID)

	log.Printf("[%s] starting metrics collection", nodeID)
	defer log.Printf("[%s] finished metrics collection", nodeID)

	// TODO(scrape all nodes, not just nxcore)
	ctx, cancel := context.WithTimeout(r.Context(), *timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ssh", "-C", "node-"+nodeID, `bash -c "kubectl get nodes | awk '/Ready/ {print \$1}' | xargs -I{} kubectl get --raw /api/v1/nodes/{}/proxy/metrics/cadvisor"`)

	b, err := cmd.Output()

	if err != nil {
		log.Printf("[%s] error during metric collection: %s", nodeID, err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Printf("[%s] error writing response: %s", nodeID, err)
		return
	}
}

func main() {
	flag.Parse()
	http.Handle("/metrics/", http.StripPrefix("/metrics/", http.HandlerFunc(handler)))
	log.Printf("starting server on :9911")
	http.ListenAndServe(":9911", nil)
}
