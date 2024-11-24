package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/KumazakiRyoha/blockchain/crypto"
	"github.com/stretchr/testify/assert"

	"github.com/KumazakiRyoha/blockchain/types"
)

func TestBlock(t *testing.T) {
	b := randomBlock(0, types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	b.Sign(privateKey)
	assert.Nil(t, b.Sign(privateKey))
	assert.Nil(t, b.Verify())
}

func TestVerifyBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privateKey))
	assert.Nil(t, b.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivateKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}


func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version: 1,
		PrevBlockHash: prevBlockHash,
		Height: height,
		Timestamp: time.Now().UnixNano(),
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)
	tx := randTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privateKey))
	return b
}