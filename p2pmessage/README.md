# 百行go代码构建p2p聊天室
- [百行go代码构建p2p聊天室](#百行go代码构建p2p聊天室)
    - [1. 上手使用](#1-上手使用)
    - [2. whisper 原理](#2-whisper-原理)
    - [3. 源码解读](#3-源码解读)
        - [3.1 参数说明](#31-参数说明)
        - [3.1 连接主节点](#31-连接主节点)
        - [3.2 我的标识](#32-我的标识)
        - [3.2 配置我的节点](#32-配置我的节点)
        - [3.3 哪个聊天室](#33-哪个聊天室)
        - [3.3 加入聊天室](#33-加入聊天室)
        - [3.4 群发消息](#34-群发消息)
        - [3.5  接收消息](#35-接收消息)
    - [4. 再次使用p2pmessage](#4-再次使用p2pmessage)

只需百行代码,就可以构建一个完整的p2p聊天室,并且消息加密,无法被追踪;并且不需要服务器,永不停机,是不是很酷.

    系统实际上基于以太坊的whisper,它本来是为以太坊上的DAPPS通信构建的,这里直接拿来做聊天室一点问题都没有.


## 1. 上手使用
先说用法,来感受一下完全匿名的P2p聊天系统.
在终端运行p2pmessage.exe,然后等待出现.
`Connected to peer,you can type message now.`,这时候你就已经连接到whisper的p2p网络中了. 有可能你需要几分钟才能成功.
你可以收到来自别人的消息,也可以发送消息给别人.
你可以同时在不同的机器上启动程序来感受一下结果.
```
2017-09-11 11:31:18 <mine>: hello
input ~Q to quit>
2017-09-11 11:33:10 [182cbbaac94313b3b96b25d9c9a0a1adea4519e9]: who am i
```
第一行是收到了我发送的消息,第二行是提示输入,第三行是收到了182cbbaac94313b3b96b25d9c9a0a1adea4519e9发来的消息.
其中182cbbaac94313b3b96b25d9c9a0a1adea4519e9是结点标识.

## 2. whisper 原理

接入whisper网络中的节点,在收到任何消息会首先验证一下工作量(可以参考bitmessage),如果没问题然后就转发.
同时也会看看是不是发送给我的,如果是就告诉用户.
至于怎么知道是不是发送给我的,有多种方式,这里只使用主题以及密码匹配.
也就是说,必须是我感兴趣的主题,同时加密的密码也是我指定的.那就可以愉快的聊天了.
当然,我没办法知道对方是谁,除非他告诉我.

## 3. 源码解读

可以到[github下载完整的源码](https://github.com/nkbai/go-ethereum-example/blob/master/p2pmessage/p2pmessage.go).

### 3.1 参数说明
总共有三个参数
-verbosity 用来打印调试信息
-topic 聊天室的主题(任意四个字节), 你必须事先知道主题才能加入,如果随便写一个,那就是你自己创建一个聊天室了.
-password 聊天室的密码, 主题密码都一致,才能进入同一个聊天室.

### 3.1 连接主节点
虽然说p2p网络没有服务器,但是必须存在知名节点,否则无从启动网络.
首先就是连接以太坊的主节点.
```go
	for _, node := range ethparams.MainnetBootnodes {
		peer := discover.MustParseNode(node)
		peers = append(peers, peer)
	}
    peer := discover.MustParseNode("enode://b89172e36cb79202dd0c0822d4238b7a7ddbefe8aa97489049c9afe68f71b10c5c9ce588ef9b5df58939f982c718c59243cc5add6cebf3321b88d752eac02626@182.254.155.208:33333")
	peers = append(peers, peer)
```
后面这个节点是我搭建的.方便国内的用户快速通信,因为基于主节点的通信可能会比较慢,延时比较长.

### 3.2 我的标识
每个节点都有自己的私钥,标识就是自己的公钥.
当然可以每次都使用相同的私钥,这里简单起见,每次都是自动生成了.
```go
    asymKeyID, err = shh.NewKeyPair()
	if err != nil {
		utils.Fatalf("Failed to generate a new key pair: %s", err)
	}

	asymKey, err = shh.GetPrivateKey(asymKeyID)
	if err != nil {
		utils.Fatalf("Failed to retrieve a new key pair: %s", err)
	}
```

### 3.2 配置我的节点
一个节点就是不停的转发符合Pow的消息,如果是我这个聊天室的消息,就告诉用户.所以节点要和其他节点进行交互,交互的节点越多,消息传播的越快.
当然这些节点数量要有一个上限,这里是80. 其中peers变量就是[3.1 连接主节点](#31-连接主节点)的主节点.
```go
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
```
### 3.3 哪个聊天室

具有相同的主题和密码的就是同一个聊天室. 
symKey关联到指定的密码,topic保存四个字节的指定主题.
```go
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
```

### 3.3 加入聊天室

我的节点可能会收到千百条各种消息,有些我能解密,有些我不能解密,但是其中只有极少一部分是我想看到的.
所以要告诉我的节点我只对这个聊天室感兴趣,如果有消息来就告诉我.
SubscribeMessage订阅指定主题和密码的消息,注意filterID,它相当于向系统订阅特定消息的句柄,后面还会用到.
```go
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
```

### 3.4 群发消息

在p2p网络中群发消息反而是最简单的,如果要点对点发消息,限制反而要多. whisper提供有发送消息的api.
主要就是构造一个合法的消息结构,主要是指定topic以及加密的秘钥,还有就是消息体(payload)就可以了,asymKey主要是为了标识ID,不是用作非对称加密.

发送消息主要是按照指定的PoW要求(比特币,莱特币,以太币等等都是类似的思路),计算hash,然后把消息发送到网络上.
```go
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

```

### 3.5  接收消息

系统实际上一直在不停的接收消息并转发,这里说的接收消息实际上就是把我们感兴趣的消息提取出来,也就是我们这个聊天室的消息.
注意这里的filterID就是 [3.3 哪个聊天室](#33-哪个聊天室)提到的,这里可以认作是聊天室的ID了.
可以看出messageLoop就是不停的轮询有没有相关聊天室的消息,目前whisper还没有实现消息推送功能.
```go
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
```

## 4. 再次使用p2pmessage

在主机1和主机2同时上运行`p2pmessage -topic ffff0000 -password 7859931` ,并等待`Connected to peer,you can type message now.`

可以看到如下截图:
主机1:
![主机1](http://images2017.cnblogs.com/blog/124391/201709/124391-20170911150628907-1497361164.jpg)
主机2:
![主机2](http://images2017.cnblogs.com/blog/124391/201709/124391-20170911150650797-1733296582.jpg)