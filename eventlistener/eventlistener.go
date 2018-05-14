package main

import (
	"log"

	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nkbai/go-ethereum-example/eventlistener/token"
)

var tokenAddr = common.HexToAddress("0xb21F4f3E7FaB55025F5A9747CAF4c4Ff18e6b407")
var toAddr = common.HexToAddress("0xc980963a77fE17026c11b97823F7160C687170D1")

func main() {
	c, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatalf("dial err %v", err)
		return
	}
	filter, err := token.NewTokenFilterer(tokenAddr, c)
	if err != nil {
		log.Fatalf("new filter err %s", err)
	}
	//1. listen any event transfer  coming...
	ch := make(chan *token.TokenTransfer, 10)
	sub, err := filter.WatchTransfer(nil, ch, nil, []common.Address{toAddr})
	if err != nil {
		log.Fatalf("watch transfer err %s", err)
	}
	go func() {
		for {
			select {
			case <-sub.Err():
				return
			case e := <-ch:
				log.Printf("new transfer event from %s to %s value=%s,at %d",
					e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
			}
		}
	}()
	//2. get history of event transfer
	history, err := filter.FilterTransfer(&bind.FilterOpts{Start: 480000}, nil, []common.Address{toAddr})
	if err != nil {
		log.Fatalf("query history logs err %s", err)
	}
	for history.Next() {
		e := history.Event
		log.Printf("%s transfer to %s value=%s, at %d", e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
	}
	//transfer to this addr
	time.Sleep(time.Minute * 3)
	log.Printf("finished..")
	sub.Unsubscribe()
}
