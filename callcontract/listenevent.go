package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-callcontract-example/mytoken"
	"log"
	"time"
)

func main() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := mytoken.NewMyToken(common.HexToAddress("0xbb4cc62817e59c68bd00e595add7e7f5b2ffacf9"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	ch, sub, err := token.EventTransferSubscribe(nil, 0, 0x1840)
	if err != nil {
		log.Fatalf("faited to subscribe event :%v", err)
	}
	go func() {
		for {
			log, ok := <-ch
			if !ok {
				break
			}
			ls, _ := json.Marshal(log)
			fmt.Printf("log:%s\n", string(ls))
		}
	}()
	time.Sleep(2 * time.Minute)
	sub.Unsubscribe()
}
