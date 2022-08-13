package main

import (
	"fmt"
	"github.com/mukund314115/gobitch/blockchain"
)

func main() {
	fmt.Println("Starting BlockChain")
	blockChain := blockchain.InitBlockChain()
	blockChain.AddBlock("First")
	blockChain.AddBlock("Second")
	blockChain.AddBlock("Third")
	for i := range blockChain.Blocks {
		fmt.Println(string(blockChain.Blocks[i].Data))
	}
}
