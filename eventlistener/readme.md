# 如何使用 geth1.8来监听合约事件

# 新功能介绍
geth1.8版本带来了新的事件处理方式,使用 abigen 可以自动生成包含合约事件监听以及过滤相关代码.
这样就不用自己去写代码解析 log.
比如:
```bash
abigen --sol token.sol --pkg token --out token.go
```
# 一个例子
通过监听谁给我转账,来说明如何使用新的接口

## 创建 Filter
只需指定合约地址即可.
```go
filter, err := token.NewTokenFilterer(tokenAddr, c)
```
## 监听将要发生的事件
这个应该放在过滤历史事件之前,因为有可能在处理历史事件过程中产生了新的事件.如果顺序错了,就会造成事件丢失.
```go
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
```
简单直观,不用去关心 log 的细节.
感兴趣的话,可以看一下 TokenTranser 结构
```go
// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
```

## 过滤历史事件
也很直观,把你感兴趣的事件范围传递进去,会返回一个 Iterator, 遍历就 ok 了.
```go
    history, err := filter.FilterTransfer(&bind.FilterOpts{Start: 480000}, nil, []common.Address{toAddr})
	for history.Next() {
		e := history.Event
		log.Printf("%s transfer to %s value=%s, at %d", e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
	}
```

# 结论
有了这些自动生成的代码以后,我们就不用费劲去理解过滤时候的 Topic 怎么设置,Log怎么解析. 直接关注我们想要的事件本身就可以了.
当然也不是没有问题,如果我关注的不是某个合约上发生了转账事件,而是所有的ERC20token, 那么该怎么写呢?
目前我是没想到怎么实现,要想这么做还是要回到老办法上.