package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/KumazakiRyoha/blockchain/crypto"
	"github.com/stretchr/testify/assert"

	"github.com/KumazakiRyoha/blockchain/types"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version: 1,
		PrevBlock: types.RandomHash(),
		Height: height,
		Timestamp: time.Now().UnixNano(),
	}
	tx := &Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{*tx})
}

func TestBlock(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	b.Sign(privateKey)
	assert.Nil(t, b.Sign(privateKey))
	assert.Nil(t, b.Verify())
}

func TestVerifyBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privateKey))
	assert.Nil(t, b.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivateKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}