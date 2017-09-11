// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// This is a simple Whisper node. It could be used as a stand-alone bootstrap node.
// Also, could be used for different test and diagnostics purposes.

package main

import (
	"bufio"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/nat"
	ethparams "github.com/ethereum/go-ethereum/params"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv5"
)

const quitCommand = "~Q"

// singletons
var (
	server *p2p.Server
	shh    *whisper.Whisper
	done   chan struct{}

	input = bufio.NewReader(os.Stdin)
)

// encryption
var (
	symKey   []byte
	asymKey  *ecdsa.PrivateKey
	topic    whisper.TopicType
	filterID string
)

// cmd arguments
var (
	argVerbosity = flag.Int("verbosity", int(log.LvlError), "log verbosity level")
	argTopic     = flag.String("topic", "44c7429f", "topic in hexadecimal format (e.g. 70a4beef)")
	argPass      = flag.String("password", "123456", "message's encryption password")
)

func main() {
	flag.Parse()
	initialize()
	run()
}

func initialize() {
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*argVerbosity), log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	done = make(chan struct{})
	var peers []*discover.Node
	var err error
	var asymKeyID string

	//connect to mainnet
	for _, node := range ethparams.MainnetBootnodes {
		peer := discover.MustParseNode(node)
		peers = append(peers, peer)
	}
	peer := discover.MustParseNode("enode://b89172e36cb79202dd0c0822d4238b7a7ddbefe8aa97489049c9afe68f71b10c5c9ce588ef9b5df58939f982c718c59243cc5add6cebf3321b88d752eac02626@182.254.155.208:33333")
	peers = append(peers, peer)
	shh = whisper.New()

	asymKeyID, err = shh.NewKeyPair()
	if err != nil {
		utils.Fatalf("Failed to generate a new key pair: %s", err)
	}

	asymKey, err = shh.GetPrivateKey(asymKeyID)
	if err != nil {
		utils.Fatalf("Failed to retrieve a new key pair: %s", err)
	}

	maxPeers := 80

	server = &p2p.Server{
		Config: p2p.Config{
			PrivateKey:     asymKey,
			MaxPeers:       maxPeers,
			Name:           common.MakeName("p2p chat group", "5.0"),
			Protocols:      shh.Protocols(),
			NAT:            nat.Any(),
			BootstrapNodes: peers,
			StaticNodes:    peers,
			TrustedNodes:   peers,
		},
	}
}

func configureNode() {
	symKeyID, err := shh.AddSymKeyFromPassword(*argPass)
	if err != nil {
		utils.Fatalf("Failed to create symmetric key: %s", err)
	}
	symKey, err = shh.GetSymKey(symKeyID)
	if err != nil {
		utils.Fatalf("Failed to save symmetric key: %s", err)
	}
	copy(topic[:], common.FromHex(*argTopic))
	fmt.Printf("Filter is configured for the topic: %x \n", topic)
}
func startServer() {
	err := server.Start()
	if err != nil {
		utils.Fatalf("Failed to start Whisper peer: %s.", err)
	}

	fmt.Println("Whisper node started,please send message after connect to other nodes")
	// first see if we can establish connection, then ask for user input
	waitForConnection(false)
	configureNode()
	SubscribeMessage()

	fmt.Printf("Please type the message. To quit type: '%s'\n", quitCommand)
}

func SubscribeMessage() {
	var err error

	filter := whisper.Filter{
		KeySym:   symKey,
		KeyAsym:  asymKey,
		Topics:   [][]byte{topic[:]},
		AllowP2P: true,
	}
	filterID, err = shh.Subscribe(&filter)
	if err != nil {
		utils.Fatalf("Failed to install filter: %s", err)
	}
}

func waitForConnection(timeout bool) {
	var cnt int
	var connected bool
	for !connected {
		time.Sleep(time.Millisecond * 500)
		connected = server.PeerCount() > 0
		if timeout {
			cnt++
			if cnt > 1000 {
				utils.Fatalf("Timeout expired, failed to connect")
			}
		}
	}

	fmt.Println("Connected to peer,you can type message now.")
}

func run() {
	startServer()
	defer server.Stop()
	shh.Start(nil)
	defer shh.Stop()
	//接收消息
	go messageLoop()
	//控制台发送消息
	sendLoop()
}

func sendLoop() {
	for {
		s := scanLine(fmt.Sprintf("input %s to quit>", quitCommand))
		if s == quitCommand {
			fmt.Println("Quit command received")
			close(done)
			break
		}
		sendMsg([]byte(s))
	}
}

func sendMsg(payload []byte) common.Hash {
	params := whisper.MessageParams{
		Src:      asymKey,
		KeySym:   symKey,
		Payload:  payload,
		Topic:    topic,
		TTL:      whisper.DefaultTTL,
		PoW:      whisper.DefaultMinimumPoW,
		WorkTime: 5,
	}

	msg, err := whisper.NewSentMessage(&params)
	if err != nil {
		utils.Fatalf("failed to create new message: %s", err)
	}
	envelope, err := msg.Wrap(&params)
	if err != nil {
		fmt.Printf("failed to seal message: %v \n", err)
		return common.Hash{}
	}

	err = shh.Send(envelope)
	if err != nil {
		fmt.Printf("failed to send message: %v \n", err)
		return common.Hash{}
	}

	return envelope.Hash()
}

func messageLoop() {
	f := shh.GetFilter(filterID)
	if f == nil {
		utils.Fatalf("filter is not installed")
	}

	ticker := time.NewTicker(time.Millisecond * 50)

	for {
		select {
		case <-ticker.C:
			messages := f.Retrieve()
			for _, msg := range messages {
				printMessageInfo(msg)
			}
		case <-done:
			return
		}
	}
}

func printMessageInfo(msg *whisper.ReceivedMessage) {
	text := string(msg.Payload)
	timestamp := time.Unix(int64(msg.Sent), 0).Format("2006-01-02 15:04:05")
	var address common.Address
	if msg.Src != nil {
		address = crypto.PubkeyToAddress(*msg.Src)
	}
	if whisper.IsPubKeyEqual(msg.Src, &asymKey.PublicKey) {
		fmt.Printf("\n%s <mine>: %s\n", timestamp, text) // message from myself
	} else {
		fmt.Printf("\n%s [%x]: %s\n", timestamp, address, text) // message from a peer
	}
}

func scanLine(prompt string) string {
	if len(prompt) > 0 {
		fmt.Print(prompt)
	}
	txt, err := input.ReadString('\n')
	if err != nil {
		utils.Fatalf("input error: %s", err)
	}
	txt = strings.TrimRight(txt, "\n\r")
	return txt
}
