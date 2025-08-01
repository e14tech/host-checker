package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Hosts []string `toml:"hosts"`
}

func main() {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	fmt.Printf("Hosts: %s\n", config.Hosts)

	fmt.Printf("Host in spot 1: %s\n", config.Hosts[1])
}
