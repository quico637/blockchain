package blockchain

type Block struct {
	Hash        [32]byte //represent the hash of this block
	Transaction *Transaction
	PrevHash    [32]byte //represent the last block hash, this allow us to link the block together
	// each block inside a blockchain references the last block that was created inside the blockchain
}
