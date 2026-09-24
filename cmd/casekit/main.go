package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/relentlessworks/casekit/internal/api"
	"github.com/relentlessworks/casekit/internal/config"
)

func main() {
	cfg := config.Load()
	handler := api.New(cfg.Secret)
	mux := handler.Routes()

	log.Printf("casekit listening on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, mux); err != nil {
		fmt.Printf("error: server failed: %v\n", err)
	}
}
