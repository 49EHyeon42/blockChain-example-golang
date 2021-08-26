package main

import (
	"github.com/49EHyeon42/blockChain-example-golang/persistence-and-CLI/blockchain"
	"github.com/49EHyeon42/blockChain-example-golang/persistence-and-CLI/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
