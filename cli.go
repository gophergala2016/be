package main

import (
	"fmt"
	"log"

	api "github.com/gophergala2016/be/insightapi"
)

func cliLatestBlocks() {
	latestBlocks, err := api.GetLatestBlocks()

	if err != nil {
		log.Fatal(err)
	}

	for _, b := range latestBlocks.Blocks[0:20] {
		fmt.Println(b.Height, b.Hash, b.Time, b.Txlength, b.PoolInfo.PoolName, b.Size)
	}
}

func cliBlock(block string) {
	b, err := api.GetBlock(block)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block #%d\n-------------\n", b.Height)
	fmt.Printf("Number Of Transactions: %d\n", len(b.Tx))
	fmt.Printf("Height: %d\n", b.Height)
	fmt.Printf("Block Reward: %f\n", b.Reward)
	fmt.Printf("Timestamp: %d\n", b.Time)
	fmt.Printf("Mined by: %s\n", b.PoolInfo.PoolName)
	fmt.Printf("Merkle Root: %s\n", b.Merkleroot)
	fmt.Printf("Difficulty: %f\n", b.Difficulty)
	fmt.Printf("Size (bytes): %d\n", b.Size)
	fmt.Printf("Nonce: %d\n", b.Nonce)

	fmt.Println("\nTransactions:\n-------------")
	for _, tx := range b.Tx {
		fmt.Println(tx)
	}
}

func cliTransaction(transaction string) {
	tx, err := api.GetTx(transaction)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction ID: %s\n", tx.Txid)
	fmt.Printf("Value Out: %f\n", tx.ValueOut)
	fmt.Printf("Value In: %f\n", tx.ValueIn)
	fmt.Printf("Size: %d\n", tx.Size)
	fmt.Printf("Fees: %f\n", tx.Fees)

	fmt.Println("\nInputs:\n-------")
	for _, i := range tx.Vin {
		fmt.Println(i.Addr, formatBtc(i.Value))
	}

	fmt.Println("\nOutputs:\n--------")
	for _, o := range tx.Vout {
		fmt.Println(o.ScriptPubKey.Addresses[0], o.Value)
	}
}

func cliAddress(address string) {
	a, err := api.GetAddr(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Address: %s\n", a.AddrStr)
	fmt.Printf("Total Received: %s\n", formatBtc(a.TotalReceived))
	fmt.Printf("Total Sent: %s\n", formatBtc(a.TotalSent))
	fmt.Printf("Final Balance: %s\n", formatBtc(a.Balance))
	fmt.Printf("Unconfirmed Balance: %s\n", formatBtc(a.UnconfirmedBalance))
	fmt.Printf("Total Transactions: %d\n", len(a.Transactions))

	fmt.Println("\nTransactions:\n-------------")
	for _, tx := range a.Transactions {
		fmt.Println(tx)
	}
}
