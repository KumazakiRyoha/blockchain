package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	//address := publicKey.Address()

	msg := []byte("hello world")
	sign, err := privateKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sign.Verify(publicKey, msg))

}

func TestPrivateKey_SignSuccess(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	//address := publicKey.Address()

	msg := []byte("hello world")
	sign, err := privateKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sign.Verify(publicKey, msg))

}

func TestPrivateKey_SignFail(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := []byte("hello world")
	sign, err := privateKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	assert.False(t, sign.Verify(otherPublicKey, msg))
	assert.False(t, sign.Verify(publicKey, []byte("hello japan")))

}
