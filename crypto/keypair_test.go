package crypto

import (
	
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestKeyPair_Sign_Verify_Success(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	b:=sig.Verify(publicKey,msg)

	 assert.True(t, b)
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("xxxxxx")))
}