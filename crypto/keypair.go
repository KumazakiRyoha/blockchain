package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/KumazakiRyoha/blockchain/types"
	"math/big"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return PrivateKey{
		key: key,
	}
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}
	return &Signature{
		r: r,
		s: s,
	}, nil
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (p PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &p.key.PublicKey,
	}
}

func (p PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(p.key, p.key.X, p.key.Y)
}

func (p PublicKey) Address() types.Address {
	h := sha256.Sum256(p.ToSlice())
	return types.NewAddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	s, r *big.Int
}

func (s Signature) Verify(publicKey PublicKey, data []byte) bool {
	return ecdsa.Verify(publicKey.key, data, s.r, s.s)

}
