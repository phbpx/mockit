package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/phbpx/mockit"
	"gopkg.in/yaml.v3"
)

func main() {
	var configPath, addr string
	flag.StringVar(&configPath, "config", "", "config file")
	flag.StringVar(&addr, "addr", ":8080", "http service address")
	flag.Parse()

	if err := run(configPath, addr); err != nil {
		log.Fatalf("run: %v", err)
	}
}

func run(configPath string, addr string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read config file: %v", err)
	}

	var config mockit.Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return fmt.Errorf("unmarshal config file: %v", err)
	}

	srv := http.Server{
		Addr:     addr,
		Handler:  mockit.NewRouter(config),
		ErrorLog: log.New(os.Stderr, "", log.LstdFlags),
	}

	log.Printf("running on %s\n", addr)
	return srv.ListenAndServe()
}
