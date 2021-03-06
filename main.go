package main

import (
	"flag"
	"fmt"
	"os"

	api "github.com/gophergala2016/be/insightapi"
)

func formatBtc(f float32) (s string) {
	return fmt.Sprintf("%.8f", f)
}

func main() {

	var (
		block       string
		transaction string
		address     string
		apiURL      string
		help        bool
		cli         bool
	)

	flag.StringVar(&block, "b", "", "Block")
	flag.StringVar(&transaction, "t", "", "Transaction")
	flag.StringVar(&address, "a", "", "Address")
	flag.StringVar(&apiURL, "api-url", api.ApiURL, "API URL")
	flag.BoolVar(&help, "h", false, "Print Help")
	flag.BoolVar(&cli, "cli", false, "CLI mode")
	flag.Parse()

	api.ApiURL = apiURL
	api.UserAgent = "be"

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if block == "" && transaction == "" && address == "" {
		if !cli {
			tuiLatestBlocks()
		} else {
			cliLatestBlocks()
		}
		os.Exit(0)
	}

	if block != "" {
		cliBlock(block)
		os.Exit(0)
	}

	if transaction != "" {
		cliTransaction(transaction)
		os.Exit(0)
	}

	if address != "" {
		cliAddress(address)
		os.Exit(0)
	}
}
