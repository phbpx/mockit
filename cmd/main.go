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

type appConfig struct {
	configSlice configSlice
	addr        string
}

func (c *appConfig) Validate() error {
	if len(c.configSlice) == 0 {
		return fmt.Errorf("no configs provided")
	}

	if c.addr == "" {
		return fmt.Errorf("no address provided")
	}

	return nil
}

type configSlice []string

func (c *configSlice) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func (c *configSlice) String() string {
	return strings.Join(*c, ", ")
}

// ============================================================================

func main() {
	cfg := newAppConfig()

	if err := run(cfg); err != nil {
		fmt.Printf("run: %v", err)
		os.Exit(1)
	}
}

func run(cfg appConfig) error {
	var merged mockit.Config
	for _, path := range cfg.configSlice {
		cfg := mockit.Config{}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read config [%s]: %v", path, err)
		}

		if err := yaml.Unmarshal(content, &cfg); err != nil {
			return fmt.Errorf("unmarshal config [%s]: %v", path, err)
		}

		merged.Endpoints = append(merged.Endpoints, cfg.Endpoints...)
	}

	srv := http.Server{
		Addr:     cfg.addr,
		Handler:  mockit.NewRouter(merged),
		ErrorLog: log.New(os.Stderr, "", log.LstdFlags),
	}

	fmt.Printf("running on %s\n", cfg.addr)
	for _, endpoint := range merged.Endpoints {
		fmt.Printf("(%s) %s\n", endpoint.Method, endpoint.URL)
	}

	return srv.ListenAndServe()
}

func newAppConfig() appConfig {
	appConfig := &appConfig{}

	flag.StringVar(&appConfig.addr, "addr", ":8080", "http service address")
	flag.Var(&appConfig.configSlice, "config", "")
	flag.Parse()

	envAddr := os.Getenv("MOCKIT_ADDR")
	if envAddr != "" {
		appConfig.addr = envAddr
	}

	envConfigs := os.Getenv("MOCKIT_CONFIG")
	if envConfigs != "" {
		appConfig.configSlice = strings.Split(strings.TrimSpace(envConfigs), ",")
	}

	return *appConfig
}
