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
	)

	flag.StringVar(&block, "b", "", "List Transactions in a Block")
	flag.StringVar(&transaction, "t", "", "Inspect Inputs and Outputs of a Transaction")
	flag.StringVar(&address, "a", "", "List Address amount and Transactions")
	flag.Parse()

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
}
