package core

import "io"

type Transaction struct{}

func (tsx *Transaction) EncodeBinary(w io.Writer) error { return nil }

func (tsx *Transaction) DecodeBinary(r io.Reader) error { return nil }