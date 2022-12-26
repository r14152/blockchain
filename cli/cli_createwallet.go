package cli

import (
	"fmt"

	"github.com/blockchain_go/wallet"
)

func (cli *CommandLineInterface) createWallet(nodeID string) {
	wallets, _ := wallet.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
