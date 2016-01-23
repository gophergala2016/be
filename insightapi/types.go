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
	Reward        float32
	PoolInfo      PoolInfo
}

type BlockList struct {
	Blocks []BlockInfo
}

type BlockInfo struct {
	Height   int
	Size     int
	Hash     string
	Time     string
	Txlength int
	PoolInfo PoolInfo
}

type Tx struct {
	Txid     string
	Version  int
	Locktime int
	Vin      []Input
	Vout     []Output
	Size     int
	ValueIn  float32
	ValueOut float32
	Fees     float32
}

type Input struct {
	Txid             string
	Vout             int
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
	BalanceSat              int
	TotalReceived           float32
	TotalReceivedSat        int
	TotalSent               float32
	TotalSentSat            int
	UnconfirmedBalance      float32
	UnconfirmedBalanceSat   int
	UnconfirmedTxApperances int
	TxApperances            int
	Transactions            []string
}

type PoolInfo struct {
	PoolName string
	Url      string
}
