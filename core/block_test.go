package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/ANNMAINAWANGARI/blockchain/crypto"
	"github.com/ANNMAINAWANGARI/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock( height uint32) *Block {
	//privKey := crypto.GeneratePrivateKey()
	header := &Header{
		Version:       1,
		PrevBlock:     types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	tx:=Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}

func TestHashBlock(t *testing.T){
	b:=randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}