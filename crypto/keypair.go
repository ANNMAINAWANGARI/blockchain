package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"math/big"

	"github.com/ANNMAINAWANGARI/blockchain/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

type Signature struct {
	r,s *big.Int
	
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		r:r,
		s:s}, nil
}

func GeneratePrivateKey() PrivateKey{
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func (k PrivateKey) PublicKey() PublicKey{
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

func (k PublicKey) Address() types.Address{
	//grab the bytes of the publickey which is toslice() then we hash it and create the address from the given bytes
	h:=sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(h[len(h)-20:])
}

func (k PublicKey) ToSlice() []byte{
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
}

func (sig Signature) String() string {
	b := append(sig.s.Bytes(), sig.r.Bytes()...)
	return hex.EncodeToString(b)
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	

	return ecdsa.Verify(pubKey.key, data, sig.r, sig.s)
}