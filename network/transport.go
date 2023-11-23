package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}
type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	Addr() NetAddr
	SendMessage(NetAddr,[]byte) error
}