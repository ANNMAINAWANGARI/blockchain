package core

import (
	"bytes"
	"testing"

	"github.com/ANNMAINAWANGARI/blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("foo"),
	}
	assert.Nil(t, tx.Sign(privKey))

	return &tx
}

func TestTxEncodeDecode(t *testing.T) {
	//create new transaction with signature
	tx := randomTxWithSignature(t)
	//create a buffer
	buf := &bytes.Buffer{}
	//encode it into the buffer
	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))

	//make new empty transaction
	txDecoded := new(Transaction)
	//decode the buffer into the new transaction
	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
	assert.Equal(t, tx, txDecoded)
}