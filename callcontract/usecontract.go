package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-callcontract-example/mytoken"
)

const key = `
{
  "address": "1a9ec3b0b807464e6d3398a59d6b0a369bf422fa",
  "crypto": {
    "cipher": "aes-128-ctr",
    "ciphertext": "a471054846fb03e3e271339204420806334d1f09d6da40605a1a152e0d8e35f3",
    "cipherparams": {
      "iv": "44c5095dc698392c55a65aae46e0b5d9"
    },
    "kdf": "scrypt",
    "kdfparams": {
      "dklen": 32,
      "n": 262144,
      "p": 1,
      "r": 8,
      "salt": "e0a5fbaecaa3e75e20bccf61ee175141f3597d3b1bae6a28fe09f3507e63545e"
    },
    "mac": "cb3f62975cf6e7dfb454c2973bdd4a59f87262956d5534cdc87fb35703364043"
  },
  "id": "e08301fb-a263-4643-9c2b-d28959f66d6a",
  "version": 3
}`

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
	toAddress := common.HexToAddress("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401")
	val, _ := token.BalanceOf(nil, toAddress)
	fmt.Printf("before transfer :%s\n", val)
	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewTransactor(strings.NewReader(key), "123")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := token.Transfer(auth, toAddress, big.NewInt(387))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatalf("tx mining error:%v\n", err)
	}
	val, _ = token.BalanceOf(nil, toAddress)
	fmt.Printf("after transfere:%s\n", val)
	fmt.Printf("tx is :%s\n", tx)
	fmt.Printf("receipt is :%s\n", receipt)
}
