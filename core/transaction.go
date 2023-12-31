package core

import (
	"fmt"

	"github.com/ANNMAINAWANGARI/blockchain/crypto"
	"github.com/ANNMAINAWANGARI/blockchain/types"
)

type Transaction struct{
	Data []byte
	PublicKey crypto.PublicKey
	Signature    *crypto.Signature
	hash types.Hash
	// firstSeen is the timestamp of when this tx is first seen locally
	firstSeen int64
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

// func (tsx *Transaction) EncodeBinary(w io.Writer) error { return nil }

// func (tsx *Transaction) DecodeBinary(r io.Reader) error { return nil }
func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig,err:=privKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.PublicKey = privKey.PublicKey()
	tx.Signature = sig
	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	//hash := tx.Hash(TxHasher{})
	if !tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}


func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}
	return tx.hash
}



func (tx *Transaction) SetFirstSeen(t int64) {
	tx.firstSeen = t
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.firstSeen
}