package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/ANNMAINAWANGARI/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func Test_Header_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:       1,
		PrevBlock: types.RandomHash(),
		Height:        10,
		Timestamp:     time.Now().UnixNano(),
		Nonce: 989394,
	}
	//make the header a binary representation into the buffer
	buf:=&bytes.Buffer{}
	assert.Nil(t,h.EncodeBinary(buf))

	//decode the header from the buffer making h equal as hDecode
	hDecode:=&Header{}
	assert.Nil(t,hDecode.DecodeBinary(buf))
	assert.Equal(t,h,hDecode)
	
	
}
func TestBlock_Encode_Decode(t *testing.T){
	b:=&Block{
		Header: Header{
			Version:       1,
		    PrevBlock: types.RandomHash(),
		    Height:        10,
		    Timestamp:     time.Now().UnixNano(),
		    Nonce: 989394,
		},
		Transactions: nil,
	}
	buf:=&bytes.Buffer{}
	assert.Nil(t,b.EncodeBinary(buf))

	bDecode:=&Block{}
	assert.Nil(t,bDecode.DecodeBinary(buf))
	assert.Equal(t,b,bDecode)
	fmt.Printf("%+v",bDecode)
}

func TestBlockHash(t *testing.T){
	b:=&Block{
		Header: Header{
			Version:       1,
		    PrevBlock: types.RandomHash(),
		    Height:        10,
		    Timestamp:     time.Now().UnixNano(),
		    
		},
		Transactions: []Transaction{},
	}
	h:=b.Hash()
	fmt.Println(h)
	assert.False(t,h.IsZero())
}