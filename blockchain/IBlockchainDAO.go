package blockchain

type IBlockchainDAO interface {
	// crud operations
	createBlockchain(b *Blockchain) bool
	getBlockchain() *Blockchain
	getPosition(b *Block) int
	getBlock(i int) *Block
	addBlock(chain *Blockchain, b *Block) bool
}
