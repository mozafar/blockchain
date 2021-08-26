package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Payload struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Block struct {
	PrevHash  []byte
	Payload   Payload
	Hash      []byte
	Timestamp int64
	Nonce     int
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(prevHash []byte, p Payload) *Block {
	b := &Block{
		PrevHash:  prevHash,
		Payload:   p,
		Timestamp: time.Now().Unix(),
	}
	b.SetHash()
	pow := NewProofOfWork(b)
	nonce, hash := pow.Run()

	b.Hash = hash[:]
	b.Nonce = nonce

	return b
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{
		b.PrevHash,
		[]byte(fmt.Sprintf("%v", b.Payload)),
		timestamp,
	}, []byte{})

	h := sha256.New()
	h.Write(headers)
	b.Hash = h.Sum(nil)
}

func (bc *Blockchain) Add(p Payload) {
	var b *Block
	if len(bc.Blocks) == 0 {
		b = NewBlock(nil, p)
	} else {
		b = NewBlock(bc.Blocks[len(bc.Blocks)-1].Hash, p)
	}

	bc.Blocks = append(bc.Blocks, b)
}
