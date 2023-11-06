package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/phbpx/mockit"
	"gopkg.in/yaml.v3"
)

type configSlice []string

func (c *configSlice) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func (c *configSlice) String() string {
	return strings.Join(*c, ", ")
}

func main() {
	var configs configSlice
	var addr string

	flag.Var(&configs, "config", "configuration file (.yml || .yaml)")
	flag.StringVar(&addr, "addr", ":8080", "http service address")
	flag.Parse()

	if err := run(configs, addr); err != nil {
		fmt.Printf("run: %v", err)
		os.Exit(1)
	}
}

func run(configs configSlice, addr string) error {
	var merged mockit.Config
	for _, path := range configs {
		cfg := mockit.Config{}
		if err := loadConfig(path, &cfg); err != nil {
			return err
		}

		merged.Endpoints = append(merged.Endpoints, cfg.Endpoints...)
	}

	handler, err := mockit.NewRouter(merged)
	if err != nil {
		return err
	}

	srv := http.Server{
		Addr:     addr,
		Handler:  handler,
		ErrorLog: log.New(os.Stderr, "", log.LstdFlags),
	}

	fmt.Printf("running on %s\n", addr)
	for _, endpoint := range merged.Endpoints {
		fmt.Printf("(%s) %s\n", endpoint.Method, endpoint.URL)
	}

	return srv.ListenAndServe()
}

func loadConfig(path string, cfg *mockit.Config) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config [%s]: %v", path, err)
	}

	if err := yaml.Unmarshal(content, cfg); err != nil {
		return fmt.Errorf("unmarshal config [%s]: %v", path, err)
	}

	return nil
}
