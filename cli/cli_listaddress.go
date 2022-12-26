package cli

import (
	"fmt"
	"log"

	"github.com/blockchain_go/wallet"
)

func (cli *CommandLineInterface) listAddresses(nodeID string) {
	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
