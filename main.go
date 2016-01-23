package main

import (
	"flag"
	"fmt"
	"os"

	api "github.com/gophergala2016/be/insightapi"
)

func listTransactions(blockHash string) {
	b, _ := api.GetBlock(blockHash)
	for _, tx := range b.Tx {
		fmt.Println(tx)
	}
}

func main() {

	var (
		block       string
		transaction string
		address     string
		apiURL      string
	)

	flag.StringVar(&block, "b", "", "Inspect Block")
	flag.StringVar(&transaction, "t", "", "Inspect Inputs and Outputs of a Transaction")
	flag.StringVar(&address, "a", "", "Get Address info")
	flag.StringVar(&apiURL, "u", api.ApiURL, "API URL")
	flag.Parse()

	api.ApiURL = apiURL

	if block == "" && transaction == "" && address == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if block != "" {
		// listTransactions(block)
		b, err := api.GetBlock(block)
		fmt.Printf("block: %#v\n", b)
		fmt.Printf("err: %s\n", err)
	}

	if transaction != "" {
		tx, err := api.GetTx(transaction)
		fmt.Printf("tx: %#v\n", tx)
		fmt.Printf("err: %s\n", err)
	}

	if address != "" {
		addr, err := api.GetAddr(address)
		fmt.Printf("tx: %#v\n", addr)
		fmt.Printf("err: %s\n", err)
	}
}
