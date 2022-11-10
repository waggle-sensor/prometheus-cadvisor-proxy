package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	nodeID, _, _ := strings.Cut(r.URL.Path, "/")
	nodeID = strings.ToLower(nodeID)

	// TODO(scrape all nodes, not just nxcore)
	ctx, cancel := context.WithCancel(r.Context())
	cmd := exec.CommandContext(ctx, "ssh", "-C", fmt.Sprintf("node-%s", nodeID), "kubectl", "get", "--raw", fmt.Sprintf("/api/v1/nodes/%s.ws-nxcore/proxy/metrics/cadvisor", nodeID))
	defer cancel()

	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		log.Printf("proxy error for %s: %s", nodeID, err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func main() {
	http.Handle("/metrics/", http.StripPrefix("/metrics/", http.HandlerFunc(handler)))
	http.ListenAndServe(":9911", nil)
}
