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

func GetBlock(blockHash string) (block Block, err error) {
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
