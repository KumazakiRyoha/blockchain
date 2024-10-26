package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/KumazakiRyoha/blockchain/crypto"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	tx := &Transaction{
		Data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)

	privKey2 := crypto.GeneratePrivateKey();
	tx.PublicKey = privKey2.PublicKey()
	assert.NotNil(t, tx.Verify())
}

func TestVerifyTransaction(t *testing.T) {
	tx := &Transaction{
		Data: []byte("foo"),
	}
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())

}