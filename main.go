package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Hosts []string `toml:"hosts"`
}

func catchSig(sigChannel <-chan os.Signal) {
	sig := <-sigChannel
	fmt.Printf("Received signal: %v (signal number %d). Shutting down\n", sig, sig.(syscall.Signal))
	os.Exit(0)
}

func main() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT)
	go catchSig(sigChannel)

	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	for {
		fmt.Printf("Hosts: %s\n", config.Hosts)
	}
}
