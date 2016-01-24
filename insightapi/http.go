package insightapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	ApiURL    = "https://insight.bitpay.com/api"
	UserAgent = "be"
)

func GetLatestBlocks() (blocklist BlockList, err error) {
	url := ApiURL + "/blocks"

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &blocklist)
	return
}

func GetBlockByHash(blockHash string) (block Block, err error) {
	url := ApiURL + "/block/" + blockHash

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &block)
	return
}

func GetBlockByHeight(blockHeight string) (block Block, err error) {
	url := ApiURL + "/block-index/" + blockHeight

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var blockIndex BlockIndex
	err = json.Unmarshal(bytes, &blockIndex)
	hash := blockIndex.BlockHash
	block, err = GetBlockByHash(hash)
	return
}

func GetTx(txId string) (tx Tx, err error) {
	url := ApiURL + "/tx/" + txId

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &tx)
	return
}

func GetAddr(addrStr string) (addr Addr, err error) {
	url := ApiURL + "/addr/" + addrStr

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &addr)
	return
}
