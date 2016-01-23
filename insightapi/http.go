package insightapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetBlock(blockHash string) (block Block, err error) {
	url := "https://insight.bitpay.com/api/block/" + blockHash

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
	url := "https://insight.bitpay.com/api/tx/" + txId

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
