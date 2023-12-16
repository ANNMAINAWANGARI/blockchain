package core

import (
	
	"crypto/sha256"
	

	"github.com/ANNMAINAWANGARI/blockchain/types"
)

//generic hasher
type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

// func (BlockHasher) Hash(b *Block) types.Hash {
// 	buf:= &bytes.Buffer{}//buffer
// 	enc:=gob.NewEncoder(buf)//encoder of the buffer
// 	if err:= enc.Encode(b.Header);err!=nil{panic(err)}
// 	h := sha256.Sum256(buf.Bytes());
// 	return types.Hash(h)
// }

func (BlockHasher) Hash(b *Header) types.Hash {
	h := sha256.Sum256(b.Bytes())
	return types.Hash(h)
}
type TxHasher struct{}


func (TxHasher) Hash(tx *Transaction) types.Hash {
	return types.Hash(sha256.Sum256(tx.Data))
}