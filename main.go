package main

import (
	"flag"
	"fmt"
	"log"
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
		help        bool
	)

	flag.StringVar(&block, "b", "", "Inspect Block")
	flag.StringVar(&transaction, "t", "", "Inspect Inputs and Outputs of a Transaction")
	flag.StringVar(&address, "a", "", "Get Address info")
	flag.StringVar(&apiURL, "u", api.ApiURL, "API URL")
	flag.BoolVar(&help, "h", false, "Print Help")
	flag.Parse()

	api.ApiURL = apiURL

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if block == "" && transaction == "" && address == "" {
		latestBlocks, err := api.GetLatestBlocks()
		if err != nil {
			log.Fatal(err)
		}
		for _, b := range latestBlocks.Blocks[0:20] {
			fmt.Println(b.Height, b.Hash, b.Time, b.Txlength, b.PoolInfo.PoolName, b.Size)
		}
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
