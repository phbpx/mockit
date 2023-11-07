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
	port        string
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
		Addr:     fmt.Sprintf("0.0.0.0:%s", cfg.port),
		Handler:  mockit.NewRouter(merged),
		ErrorLog: log.New(os.Stderr, "", log.LstdFlags),
	}

	fmt.Printf("running on %s\n", cfg.port)
	for _, endpoint := range merged.Endpoints {
		fmt.Printf("(%s) %s\n", endpoint.Method, endpoint.URL)
	}

	return srv.ListenAndServe()
}

func newAppConfig() appConfig {
	appConfig := &appConfig{}

	flag.StringVar(&appConfig.port, "port", "8080", "http port to listen on")
	flag.Var(&appConfig.configSlice, "config", "config file path")
	flag.Parse()

	envPort := os.Getenv("MOCKIT_PORT")
	if envPort != "" {
		appConfig.port = envPort
	}

	envConfigs := os.Getenv("MOCKIT_CONFIG")
	if envConfigs != "" {
		appConfig.configSlice = strings.Split(strings.TrimSpace(envConfigs), ",")
	}

	return *appConfig
}
