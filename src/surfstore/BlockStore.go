package surfstore

import (
	"fmt"
	"crypto/sha256"
    "encoding/hex"
)

type BlockStore struct {
	BlockMap map[string]Block
}

func (bs *BlockStore) GetBlock(blockHash string, blockData *Block) error {
	//panic("todo")
	if val, ok := bs.BlockMap[blockHash]; ok{
		*blockData = val
		return nil
	}
	//*blockData = nil
	return fmt.Errorf("SERVER ERROR: Block with specified hash does not exist in blockstore")
	//return err.Error()
}

func (bs *BlockStore) PutBlock(block Block, succ *bool) error {
	//panic("todo")
	blockHash, err := bs.ComputeHash(block.BlockData)
	if err != nil {
		return err
	}
	bs.BlockMap[blockHash] = block
	return nil
}

func (bs *BlockStore) HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error {
	//panic("todo")
	for _,blockHash := range blockHashesIn {
		if _, ok := bs.BlockMap[blockHash]; ok{
			*blockHashesOut = append(*blockHashesOut,blockHash)
		}
	}
	return nil
}

//Compute block hash using sha-256
func (bs *BlockStore) ComputeHash(block []byte) (string,error){
   	hash := sha256.Sum256(block)
    sha256_hash := hex.EncodeToString(hash[:])
    return sha256_hash, nil
}

// This line guarantees all method for BlockStore are implemented
var _ BlockStoreInterface = new(BlockStore)
