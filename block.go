package main

import "time"

type Block struct {
	Index        int
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Data         string
}

// NewBlock creates a new block
func NewBlock(
	index int,
	hash string,
	previousHash string,
	timestamp time.Time,
	data string,
) *Block {
	return &Block{
		Index:        index,
		Hash:         hash,
		PreviousHash: previousHash,
		Timestamp:    timestamp,
		Data:         data,
	}
}
