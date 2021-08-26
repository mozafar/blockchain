package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getFakePayload() Payload {
	return Payload{
		Sender:   "A",
		Receiver: "B",
		Amount:   100,
	}
}

func TestGenesisBlockHasNoPrevHash(t *testing.T) {
	var bc Blockchain
	bc.Add(getFakePayload())

	assert.Nil(t, bc.Blocks[0].PrevHash)
}

func TestSecondBlockHasPrevHash(t *testing.T) {
	var bc Blockchain

	bc.Add(getFakePayload())
	bc.Add(getFakePayload())

	assert.NotNil(t, bc.Blocks[1].PrevHash)
	assert.Equal(t, bc.Blocks[0].Hash, bc.Blocks[1].PrevHash)
}

func TestAddBlock(t *testing.T) {
	var bc Blockchain
	bc.Add(getFakePayload())

	assert.Len(t, bc.Blocks, 1)
}

func TestValidateProofOfWork(t *testing.T) {
	b := NewBlock(nil, getFakePayload())

	pow := NewProofOfWork(b)
	assert.True(t, pow.Validate())
}
