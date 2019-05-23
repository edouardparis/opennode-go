package main

import "flag"

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "listen", ":8080", "server listen address")
	flag.Parse()

	NewServer(listenAddr).Run()
}
