package insight

// Go implementation of the Insight API https://insight.is

type Block struct {
	Hash          string      `json:"hash"`
	Confirmations int         `json:"confirmations"`
	Size          int         `json:"size"`
	Height        int         `json:"height"`
	Version       int         `json:"version"`
	Merkleroot    string      `json:"merkleroot"`
	Tx            []InsightTx `json:"tx:omitempty"`
	Time          int         `json:"time"`
	Nonce         int         `json:"nonce"`
	Bits          string      `json:"bits"`
	Difficulty    int         `json:"difficulty"`
	Chainwork     string      `json:"chainwork"`
	Nextblockhash string      `json:"nextblockhash"`
	IsMainChain   bool        `json:"isMainChain"`
}

type Tx struct {
	Txid     string   `json:"txid"`
	Version  int      `json:"version"`
	Locktime int      `json:"locktime"`
	Vin      []Input  `json:"vin"`
	Vout     []Output `json:"vout"`
	ValueOut float32  `json:"valueOut"`
	Size     int      `json:"size"`
	ValueIn  float32  `json:"valueIn"`
	Fees     float32  `json:"fees"`
}

type Input struct {
	Txid             string   `json:"txid"`
	Vout             int      `json:"vout"`
	ScriptSig        []string `json:"scriptSig"`
	Sequence         int      `json:"sequence"`
	N                int      `json:"n"`
	Addr             string   `json:"addr"`
	ValueSat         float32  `json:"valueSat"`
	Value            float32  `json:"value"`
	DoubleSpentTxID  nil      `json:"doubleSpentTxID"`
	IsConfirmed      bool     `json:"isConfirmed"`
	Confirmations    int      `json:"confirmations"`
	UnconfirmedInput bool     `json:"unconfirmedInput"`
}

type Output struct {
	Value        string       `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}
