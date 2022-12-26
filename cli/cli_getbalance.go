package cli

import (
	"fmt"
	"log"

	"github.com/blockchain_go/blockchain"
	"github.com/blockchain_go/utils"
	"github.com/blockchain_go/utxo"
	"github.com/blockchain_go/wallet"
)

func (cli *CommandLineInterface) getBalance(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.NewBlockchain(nodeID)
	utxoSet := utxo.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	utxosList := utxoSet.FindUTXO(pubKeyHash)

	for _, out := range utxosList {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
