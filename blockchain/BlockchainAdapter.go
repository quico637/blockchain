package blockchain

import (
	"bufio"
	"encoding/json"
	"os"
)

type BlockchainAdapter struct {
	File     string
	Sep      string // separator
	currLine int
	w        *bufio.Writer
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (db BlockchainAdapter) createBlockchain(b *Blockchain) bool {

	f, err := os.OpenFile(DB_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	out, err := json.Marshal(b)
	check(err)

	if _, err = f.WriteString("BLOCKCHAIN: " + string(out)); err != nil {
		panic(err)
	}

	if _, err = f.WriteString("\n" + db.Sep); err != nil {
		panic(err)
	}

	db.currLine = 3

	return true
}

func (db BlockchainAdapter) getBlockchain() *Blockchain {
	return nil
}

func (db BlockchainAdapter) getPosition(b *Block) int {
	return 1
}

func (db BlockchainAdapter) getBlock(i int) *Block {
	return nil
}

func (db BlockchainAdapter) addBlock(chain *Blockchain, b *Block) bool {
	f, err := os.OpenFile(DB_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	out, err := json.Marshal(b)
	check(err)

	if _, err = f.WriteString("BLOCK: " + string(out)); err != nil {
		panic(err)
	}

	if _, err = f.WriteString("\n" + db.Sep); err != nil {
		panic(err)
	}
	db.currLine += 2

	return true
}
