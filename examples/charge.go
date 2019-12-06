package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/edouardparis/opennode-go/opennode"
	"github.com/edouardparis/opennode-go/opennode/client"
)

// go run charge.go --amount=<amt> --key<key>
func main() {
	key := flag.String("key", "", "opennode api key")
	amount := flag.Int("amount", 10, "charge amount")
	mainnet := flag.Bool("mainnet", false, "use mainnet")
	flag.Parse()
	if *key == "" {
		log.Fatal("no api key provided")
	}

	env := opennode.Development
	if *mainnet {
		env = opennode.Production
	}

	clt := client.New(*key, env)
	charge, err := clt.CreateCharge(&opennode.ChargePayload{
		Amount: int64(*amount),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("id: %s\n", charge.ID)
	fmt.Println("invoice:", charge.LightningInvoice.PayReq)
}
