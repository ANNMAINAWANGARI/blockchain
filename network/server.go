package network

import (
	"fmt"
	"time"

	"github.com/ANNMAINAWANGARI/blockchain/core"
	"github.com/ANNMAINAWANGARI/blockchain/crypto"
	"github.com/sirupsen/logrus"
)

type ServerOpts struct {
	Transports []Transport
	BlockTime  time.Duration
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts
	blockTime   time.Duration
	rpcCh  chan RPC
	memPool     *TxPool
	isValidator bool
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		blockTime:   opts.BlockTime,
		memPool: NewTxPool(),
		rpcCh:      make(chan RPC),
		isValidator: opts.PrivateKey != nil,//if we have a privatekey then we are a validator
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker:= time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n",rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do stuff every xseconds")
		}
	}

	fmt.Printf("Server shutdown")
}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("transaction already in mempool")

		return nil
	}

	logrus.WithFields(logrus.Fields{
		"hash": hash,
	}).Info("adding new tx to the mempool")

	return s.memPool.Add(tx)
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)

	}
}
