package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockChain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestBlockchain(t *testing.T) {
	bc, err := NewBlockChain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	for i := 0; i < 1000; i++ {
		b := randomBlock(uint32(i + 1))
		assert.Nil(t, bc.AddBlock(b))
	}
	assert.Equal(t, bc.Height(), uint32(1000))
	assert.Equal(t, len(bc.headers), 1000+1)
	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}

