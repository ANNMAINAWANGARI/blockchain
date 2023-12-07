package core

import (
	"fmt"
	"io"

	"github.com/ANNMAINAWANGARI/blockchain/crypto"
)

type Transaction struct{
	Data []byte
	PublicKey crypto.PublicKey
	Signature    *crypto.Signature
}

func (tsx *Transaction) EncodeBinary(w io.Writer) error { return nil }

func (tsx *Transaction) DecodeBinary(r io.Reader) error { return nil }

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