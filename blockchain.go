package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Blockchain struct {
	ID      string
	name    string
	version float32
	blocks  []Block
}

func (b *Blockchain) Init() {
	b.ID = "initial blockchain"
	b.name = "Cryptocoin"
	b.version = 0.1
	b.blocks = append(b.blocks, b.genesisBlock())
	for i := 0; i < 10; i++ {
		nextBlock := b.createNextBlock("")
		b.blocks = append(b.blocks, nextBlock)
		fmt.Printf("%#v", nextBlock)
		time.Sleep(time.Second * 1)
	}
}

func (b *Blockchain) genesisBlock() Block {
	return Block{
		Index:        0,
		Hash:         "00000000000000000000000000000000",
		PreviousHash: "",
		Timestamp:    time.Now(),
		Data:         "Genesis Block",
	}
}

func (b *Blockchain) createNextBlock(data string) Block {
	previousBlock := b.blocks[len(b.blocks)-1]
	nextIndex := previousBlock.Index + 1
	nextTimestamp := time.Now()
	nextHash := b.calculateHash(nextIndex, previousBlock.Hash, nextTimestamp, data)

	return Block{
		Index:        nextIndex,
		Hash:         nextHash,
		PreviousHash: previousBlock.Hash,
		Timestamp:    time.Now(),
		Data:         data,
	}
}

func (b *Blockchain) calculateHash(index int, hash string, timestamp time.Time, data string) string {
	newHash := sha256.Sum256([]byte(fmt.Sprintf("%d%s%s%s", index, hash, timestamp, data)))
	return hex.EncodeToString(newHash[:])
}
