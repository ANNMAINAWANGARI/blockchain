package network

type NetAddr string


type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	Addr() NetAddr
	SendMessage(NetAddr,[]byte) error
}