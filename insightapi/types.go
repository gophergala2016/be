package insightapi

// Go implementation of the Insight API https://insight.is

type Block struct {
	Hash          string
	Confirmations int
	Size          int
	Height        int
	Version       int
	Merkleroot    string
	Tx            []string
	Time          int
	Nonce         int
	Bits          string
	Difficulty    float64
	Chainwork     string
	Nextblockhash string
	IsMainChain   bool
	PoolInfo      PoolInfo
}

type Tx struct {
	Txid     string
	Version  int
	Locktime int
	Vin      []Input
	Vout     []Output
	ValueOut float32
	Size     int
	ValueIn  float32
	Fees     float32
}

type Input struct {
	Txid             string
	Vout             int
	ScriptSig        []string
	Sequence         int
	N                int
	Addr             string
	ValueSat         float32
	Value            float32
	DoubleSpentTxID  string
	IsConfirmed      bool
	Confirmations    int
	UnconfirmedInput bool
}

type Output struct {
	Value        string
	N            int
	ScriptPubKey ScriptPubKey
}

type ScriptPubKey struct {
	Asm       string
	Hex       string
	ReqSigs   int
	Type      string
	Addresses []string
}

type Addr struct {
	AddrStr                 string
	Balance                 float32
	BalanceSat              float32
	TotalReceived           float32
	TotalReceivedSat        float32
	TotalSent               float32
	TotalSentSat            float32
	UnconfirmedBalance      int
	UnconfirmedBalanceSat   int
	UnconfirmedTxApperances int
	TxApperances            int
	Transactions            []string
}

type PoolInfo struct {
	PoolName string
	Url      string
}
