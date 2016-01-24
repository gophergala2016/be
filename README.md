# be (blockchain explorer)

**be** is a Bitcoin Blockchain explorer for the terminal that uses the [Insight API](https://insight.is/) to get data from the Blockchain. By default it uses the Insight Live URL but in case you're running your own node, the URL can be defined by `--api-url <url>`

Both TUI and CLI modes are implemented. TUI is used by default, you can switch to CLI mode using the `--cli` parameter.

## What is a blockchain explorer?

In the Bitcoin network, all new data about bitcoin addresses and transactions
is periodically archived into immutable blocks, which build on top of the
previous block that was archived. This distributed chain of blocks is known as
the blockchain. You can think of it as a Git repository for money.

A [blockchain explorer](https://en.bitcoin.it/wiki/Block_chain_browser) is a
tool that allows you to inspect the status of the blockchain and obtain
information about any block, address or transaction that took place in the
Bitcoin network. **be** is a blockchain explorer for the terminal, which uses
Insight's API; [Insight](https://insight.bitpay.com/) is a web blockchain
explorer.

## Use

### Watch latest blocks mined

```
$ be
```

Use the arrow keys to move between blocks, press Enter to examine a block (as
in `be -b <height>`) and press Q to exit.

![be screenshot](https://raw.github.com/gophergala2016/be/master/img/be_screenshot.jpg)

### Watch latest blocks mined (CLI mode)

```
$ be --cli
#394785 (2016-01-24 12:31:12) 745txs 503Kb [Discus Fish]
#394784 (2016-01-24 12:26:27) 1501txs 971Kb
#394783 (2016-01-24 12:23:41) 1983txs 731Kb [Slush]
#394782 (2016-01-24 12:13:07) 2847txs 965Kb [BitMinter]
#394781 (2016-01-24 11:48:05) 128txs 121Kb [AntMiner]
#394780 (2016-01-24 11:46:35) 1053txs 707Kb
#394779 (2016-01-24 11:43:57) 2715txs 976Kb [Discus Fish]
#394778 (2016-01-24 11:15:13) 940txs 471Kb
#394777 (2016-01-24 11:08:22) 248txs 121Kb
#394776 (2016-01-24 11:07:02) 185txs 142Kb [Discus Fish]
#394775 (2016-01-24 11:05:54) 824txs 889Kb
#394774 (2016-01-24 11:00:20) 450txs 378Kb
#394773 (2016-01-24 10:58:11) 1699txs 677Kb
#394772 (2016-01-24 10:51:05) 1446txs 926Kb
#394771 (2016-01-24 10:41:35) 2651txs 926Kb
#394770 (2016-01-24 10:15:20) 965txs 441Kb
#394769 (2016-01-24 10:07:50) 1301txs 639Kb
#394768 (2016-01-24 09:56:42) 389txs 207Kb [AntMiner]
#394767 (2016-01-24 09:53:51) 844txs 536Kb
#394766 (2016-01-24 09:48:33) 353txs 216Kb
```

### Address (CLI mode)

Get address information, including balance, amount sent, amount received and transaction list.

```
$ be --cli -a 18KfSXSHJVHK1CMS5yvbfaAfW318K9NCVd
Address: 18KfSXSHJVHK1CMS5yvbfaAfW318K9NCVd
Total Received: 9.53021049
Total Sent: 9.53021049
Final Balance: 0.00000000
Unconfirmed Balance: 0.00000000
Total Transactions: 2

Transactions:
-------------
b8f03e1ae8b3b070da93b27ec1064f75c1803116e0efcfc5b5dbfa405b3059c6
6167cf5ed33449586d438c7bd3d6e05410749b483453173bb6b8cbf04b178820

```

### Transaction (CLI mode)

Inspect a transaction, including inputs, outputs, confirmations, block hash, value in and value out.

```
$ be --cli -t 43699b5a4a9f2041bf8d3307f2944e0c25edbf800041a6c5df3511f337e4a61b
Transaction ID: 43699b5a4a9f2041bf8d3307f2944e0c25edbf800041a6c5df3511f337e4a61b
Received Time: 2016-01-11 21:10:16
Mined Time: 2016-01-11 21:10:16
Value Out: 0.063286
Value In: 0.063386
Size (bytes): 226
Fees: 0.000100
Mined in block: 000000000000000007fd9bb20f163ffa52d61991b630e7102394938ad776e200
Confirmations: 2009

Inputs:
-------
1NBqPcBWSpJs6dN9t2cCXj6k1ySngj4TKp 0.06338620

Outputs:
--------
15G6L7HTRzwvxrqWFgkx18DgKZyt3VcJgu 0.02200000
1NBqPcBWSpJs6dN9t2cCXj6k1ySngj4TKp 0.04128620

```

### Block (CLI mode)

Inspect blocks, including number of confirmations, list of transactions, size, block reward, height and miner. Takes either block hash or height as arguments.

```
$ be --cli -b 0000000000000000016491fbaabd1db0a89a4a281c4a14c27e3ec56f20c8b2bb | head -n 20
Block #394654
-------------
Confirmations: 132
Number Of Transactions: 1419
Height: 394654
Block Reward: 25.000000
Timestamp: 2016-01-23 16:47:26
Mined by: AntMiner
Merkle Root: 54065db739429e50d887968cc130d34cda9b2a9571d5d38e58ba58436352f8de
Difficulty: 113354299801.471130
Size (Kilobytes): 911
Nonce: 3611731857

Transactions:
-------------
4d44b0b82e3b376d3d10fb5aef7e9eebba773edec87246fc517f5edef50fd7b2
559f31cf33dc16eb9c3c2d1f15728764aa394cb6e7b1eaecf6c188f2549e87c6
6238d611e4d2a69b49ae06f96bc2b7e5aa52469220b0e1ec4b86b20a60b34acf
bbd5909dd1bbd4494d89aa82f5799a85ea5704634c001fa7516bf9b04b44a117
93d23a3b566fa81f59ccdb5d28068ec6a40a81a48fae7235997e9e58f574e00d
```

## Install

Install in your `$GOPATH/bin`:

```
go get -u github.com/gophergala2016/be
```

Using Docker:

```
$ make container
$ docker run --rm -it be
```

## TODO

- [x] CLI mode for `-b`, `-t` and `-a`
- [x] Parametrize API URL with `--api-url`
- [x] TUI mode for Last Blocks
- [ ] TUI mode for `-b`
- [ ] TUI mode for `-t`
- [ ] TUI mode for `-a`
- [ ] Follow mode for CLI `-f [blocks|transactions]`
