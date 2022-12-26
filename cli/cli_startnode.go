package cli

import (
	"fmt"
	"log"

	"github.com/blockchain_go/server"
	"github.com/blockchain_go/wallet"
)

func (cli *CommandLineInterface) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	server.StartServer(nodeID, minerAddress)
}
