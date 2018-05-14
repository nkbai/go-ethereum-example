package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-callcontract-example/mytoken"
	"log"
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

	contractName, err := token.Name(nil)
	if err != nil {
		log.Fatalf("query name err:%v", err)
	}
	fmt.Printf("MyToken Name is:%s\n", contractName)
	balance, err := token.BalanceOf(nil, common.HexToAddress("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401"))
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf("0x8c1b2e9e838e2bf510ec7ff49cc607b718ce8401's balance is %s\n", balance)
}
