package bitcoind

// Represents a block
type Block struct {
	// The block hash
	Hash string `json:"hash"`

	// The Block Flag "proof-of-work" or "proof-of-stake"
	Flags string `json:"flags"`

	// The number of confirmations
	Confirmations uint64 `json:"confirmations"`

	// The block size
	Size uint64 `json:"size"`

	StrippedSize uint64 `json:"strippedsize"`

	Weight uint64 `json:"weight"`

	// The block height or index
	Height uint64 `json:"height"`

	// The block version
	Version uint32 `json:"version"`

	VersionHex string `json:"versionHex"`

	// The merkle root
	Merkleroot string `json:"merkleroot"`

	// Slice on transaction ids
	Tx []string `json:"tx"`

	// The block time in seconds since epoch (Jan 1 1970 GMT)
	Time int64 `json:"time"`

	// The median block time in seconds since epoch Jan 1 1970 GMT
	MedianTime int64 `json:"mediantime"`

	//The current mint value
	Mint float32 `json:"mint"`

	// The nonce
	Nonce uint64 `json:"nonce"`

	// The bits
	Bits string `json:"bits"`

	// The difficulty
	Difficulty float64 `json:"difficulty"`

	// Total amount of work in active chain, in hexadecimal
	Chainwork string `json:"chainwork,omitempty"`

	// The number of transactions in the block
	NTx int32 `json:"nTx"`

	// The hash of the previous block
	Previousblockhash string `json:"previousblockhash"`

	// The hash of the next block
	Nextblockhash string `json:"nextblockhash"`
}
