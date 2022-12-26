package blockchain

import (
	"log"

	"github.com/boltdb/bolt"

	"github.com/blockchain_go/block"
	"github.com/blockchain_go/utils"
)

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *block.Block {
	var blocks *block.Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utils.BlocksBucket))
		encodedBlock := b.Get(i.currentHash)
		blocks = block.DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = blocks.PrevBlockHash

	return blocks
}
