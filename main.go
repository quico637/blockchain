package main

import (
	"fmt"

	"github.com/quico637/blockchain/blockchain"
)

func main() {

	fmt.Println("hello world!")
	bchain := blockchain.CreateEmptyBlockchain()
	fmt.Println(*bchain)
	t := blockchain.CreateTransaction(22, "33323P", "988998S", 43, "ES64-2222", "ES65-33333")

	bchain.AddTransaction(t)
	fmt.Println(*bchain)
}
