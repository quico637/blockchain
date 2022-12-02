package blockchain

import (
	"crypto/sha256"
)

const DB_FILE = "./blockchain/blockchain.db"

type Blockchain struct {
	GenesisBlock *Block
	LastBlock    *Block
	Blocks       []*Block
	LastHash     [32]byte
	Database     string
	n            int
	Db           IBlockchainDAO
}

func CreateBlockchain(b *Block) *Blockchain {
	blockchain := &Blockchain{b, b, []*Block{b}, b.Hash, DB_FILE, 1, nil}
	blockchain.Db = BlockchainAdapter{File: DB_FILE, Sep: "---\n"}
	blockchain.Db.createBlockchain(blockchain)
	return blockchain
}

func CreateEmptyBlockchain() *Blockchain {
	genesisMessage := "Genesis Message mentiendono"
	firstTransaction := &Transaction{1, "52041159W", "52041159W", 100, "ES64-0001", "ES64-0001"}
	prevHash := sha256.Sum256([]byte(genesisMessage))
	hash := sha256.Sum256([]byte(firstTransaction.ToString()))
	block := &Block{hash, firstTransaction, prevHash}
	return CreateBlockchain(block)
}

func (b *Blockchain) AddTransaction(t *Transaction) bool {
	hash := sha256.Sum256([]byte(t.ToString()))
	block := &Block{hash, t, b.LastHash}
	b.LastBlock = block
	b.LastHash = hash
	b.Blocks = append(b.Blocks, block)
	b.n++

	// compprobar si se puede meter la transaccion
	b.Db.addBlock(b, block)
	return true
}
