package cli

import (
	"fmt"
	"log"

	"github.com/blockchain_go/blockchain"
	"github.com/blockchain_go/utxo"
	"github.com/blockchain_go/wallet"
)

func (cli *CommandLineInterface) createBlockchain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := utxo.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
