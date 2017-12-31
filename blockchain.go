package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type Network interface {
	Connect()
}

type Blockchain struct {
	ID           string
	name         string
	version      float32
	blocks       []Block
	Network      Network
	blockStorage *bolt.DB
}

func (b *Blockchain) Init() {
	b.ID = "initial blockchain"
	b.name = "Cryptocoin"
	b.version = 0.1

	b.Network.Connect()
	b.bootstrap()

	b.blocks = append(b.blocks, b.genesisBlock())
	for i := 0; i < 10; i++ {
		nextBlock := b.createNextBlock("")
		b.blocks = append(b.blocks, nextBlock)
	}
}

func (b *Blockchain) bootstrap() {
	var err error

	db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalln("Error connecting to blockchain.")
	}
	defer db.Close()

	b.blockStorage = db

	tx, err := b.blockStorage.Begin(true)
	if err != nil {
		log.Fatalln("Error opening connection to blockchain.")
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("blocks"))
	if err != nil {
		log.Fatalln("Failed to initialize blockchain")
	}

	_, err = tx.CreateBucketIfNotExists([]byte("blocks"))
	if err != nil {
		log.Fatalln("Failed to initialize blockchain")
	}

	if err := tx.Commit(); err != nil {
		log.Fatalln("Failed to initialize blockchain")
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
