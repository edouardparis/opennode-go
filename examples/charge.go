package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/edouardparis/opennode-go/opennode"
)

func main() {
	key := flag.String("key", "", "opennode api key")
	mainnet := flag.Bool("mainnet", false, "use mainnet")
	flag.Parse()
	if *key == "" {
		log.Fatal("no api key provided")
	}

	env := opennode.Development
	if *mainnet {
		env = opennode.Production
	}
	fmt.Println(*key)
	fmt.Println(env)
}
