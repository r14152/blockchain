package cli

import (
	"fmt"

	"github.com/blockchain_go/blockchain"
	"github.com/blockchain_go/utxo"
)

func (cli *CommandLineInterface) reindexUTXO(nodeID string) {
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := utxo.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
