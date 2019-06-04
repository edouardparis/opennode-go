package main

import (
	"flag"
	"log"

	"github.com/edouardparis/opennode-go/opennode"
)

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "listen", ":8080", "server listen address")

	var apiKey string
	flag.StringVar(&apiKey, "api_key", "", "opennode api key")

	var mainnet bool
	flag.BoolVar(&mainnet, "mainnet", false, "use mainnet network")

	var help bool
	flag.BoolVar(&help, "help", false, "print help")

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if mainnet {
		log.Println("network used is mainnet")
	}

	if apiKey == "" {
		log.Fatal("opennode api_key is missing")
	}

	client := opennode.NewClient(&opennode.Config{APIKey: apiKey, Testnet: true})

	NewServer(listenAddr, NewHandler(client)).Run()
}
