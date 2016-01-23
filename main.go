package main

import (
	"flag"
	"os"
)

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
}
