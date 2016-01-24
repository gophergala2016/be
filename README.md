# be (blockchain explorer)

**be** is a Bitcoin Blockchain explorer for the terminal that uses the [Insight API](https://insight.is/) to get data from the Blockchain. By default it uses the Insight Live URL but in case you're running your own node, the URL can be defined by `--api-url <url>`

Both TUI and CLI modes are implemented. TUI is used by default, you can switch to CLI mode using the `--cli` parameter.

## Examples

### Watch latest blocks mined

```
$ be
```

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
$ be --cli -a 14gDc4MDLXAQJsfjZjptCpyLNnxVNhqdsP
Address: 14gDc4MDLXAQJsfjZjptCpyLNnxVNhqdsP
Total Received: 0.00267301
Total Sent: 0.00267301
Final Balance: 0.00000000
Unconfirmed Balance: 0.00000000
Total Transactions: 28

Transactions:
-------------
1755fddb10b00008e12e2b09883151fa8159e31a8b3f8077f3787c4644e111f3
31d78c1516a20440bfb00364022e2919f28f44e90fbd7d95ea8458f3d3faa86f
714c57958d2356fd26226933d431fef2e47847ae72da4ed51d3c10e892652e24
4d3b3ac13d3ff1ff2b56781af4be00f2be3f9d735389f2b0913d9dd16e2e28b7
14a2fc11f41c8c78f8df38f922d9adf3aa9ec5ecc7dc98d82455197b38f5f452
609c2b6a428f0755b074abfbd37b1f10f02e4ad41f1801fdfa796325044c4bea
9a12437617a218dd4407fbff4e7982a0924108e4db60882bfbc1faff29263ce2
d97c2f399da42023403500eaa83f2c48f3d7c08f88b50170c755ddafd1bd3e2d
a3f2939f72a2ec61bbc56c1d400baf1242fb925d50cd5bfe4bc944261b86905d
3a017dfba4538c70139fe17018bbc7260d030040a14db2f1672f885bf6d6a0ee
56ac2d8093bc4763754f611b75e2152871937ca337c75deb6999c2046ed453e0
f8a279f8ad5019432a68379ece9e27c87e92908112f048bd0cb6179bf522d2f0
1d1dfe7477c43de8b62a5e74a08d0e648b8968b807cd435856d9cbe86fc4d5b6
9b87bd98f79d0e6d7e7197b01e8b581bf719e5fae15796ae6a9abfbefb80c35e
3c659e29943688e801d1cbe267991941348b994e5ed29d2a6a90f7b38908d36e
bf4b65f14c7aef7d598bbe2a7ca1437b24891d6179081b931868718acf6bb679
b8cb7c5d445193eda77110a75002cdeef4f690cc353696fd06a6346cdf8875bf
f62e4ce53868e42260e920286a08afd5757bc3da03f9c0b9062ba3ee1c5d2ceb
5dfe6172f552499d132e5fc9b38b2353bdeba52501c2448ba4ecb95c85d9279c
479d6fdb271c3b20fb7603baa0abfd6354e626f83097e1e0dd7ba0870095dd0b
2c3ef94ec0435bbbad39c76ca70b1cb6ef08df5fd1bf94c4f6fe11395b37766b
f0d653cf2984d7e5508a7f66d342fe67d25df9ac28a5e2d32ff180ff8d81c116
2d2a16070d461822d581b4544c34cdf077758e15e88c0854483b441e9e052e14
d3e44edf89e21d0591fbfe53570cd7f6c9c43f28dc4b1c4d426c7a4fd6306795
8ae7cb3cc18a760532fb31c10ce5515a173cabc666e0559aa0f099955b2871f5
27d9745de97f8bfe5368cb1d844e989fe5a10965844c03a73aaa91727e05fc1d
ccb0d95261fac76e78e5cc4257e9a4724f8d65a0933430a66e26fa998ad7590a
e2ea68782c0667d5e538a84e5534d2b0ef53310ec409a58e2e9bdb4c2ef5e848
```

### Transaction (CLI mode)

Inspect a transaction, including inputs, outputs, confirmations, block hash, value in and value out.

```
$ be --cli -t 4b8f9d2f4a53eb2c8de6ea54ef538edef19031792c4e1bcf003bea4c01d3a1da
Transaction ID: 4b8f9d2f4a53eb2c8de6ea54ef538edef19031792c4e1bcf003bea4c01d3a1da
Received Time: 2016-01-23 17:04:16
Mined Time: 2016-01-23 17:04:16
Value Out: 0.001620
Value In: 0.001727
Size (bytes): 1958
Fees: 0.000107
Mined in block: 00000000000000000602b10b043915f743ceb4ccb60a9bba34dbe3b7a228cbc7
Confirmations: 128

Inputs:
-------
1Ap15ixzkDfmCfVtPcAqNSVgSESs7MeDYj 0.00016683
1Ap7moFDh2WSkm83pe6c2hs8hVfe6dwUd2 0.00007628
1Ap82bWoiwH9fKC7PHUtycz2YqsThYta1J 0.00014891
1Apba5wcyS39EK3mwoJFmGn2cP7tugiCQY 0.00013447
1ApcaU3wbcW7f7DgVh6estnqGBRzikiSy2 0.00015237
1APcb47bVN2GHo7PbBHwY1yR1mpLWqoRkH 0.00019356
1ApcGc5WdwcfFiEpQhY9Cgcryyo33erhFL 0.00020502
1ApDhG7RWjhHHFraFs3FRm2cUh1gHL8xE2 0.00007163
1ApDjdtV6e8Dp2XCAcsM1v1tx1rUneh64K 0.00009384
1APdYefAkw2eSrE3pbJkMfS3noYHhw7S6F 0.00014644
1aPE25iaGSznNtHCpYQn8oie5E3UvVL8U 0.00010228
1Aphbh9o5AcsUnSdeKopSpNGhTv8QY6Mza 0.00014092
1APHRdwZWXXUmKrb8rzuiMQGGR1yo4FyQh 0.00009461

Outputs:
--------
1H6YbozjMSaARR2AHTSkNyC6S63S4LD2JB 0.00162000
```

### Block (CLI mode)

Inspect blocks, including number of confirmations, list of transactions, size, block reward, height and miner.

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

## TODO

- [x] CLI mode for `-b`, `-t` and `-a`
- [x] Parametrize API URL with `--api-url`
- [x] TUI mode for Last Blocks
- [ ] TUI mode for `-b`
- [ ] TUI mode for `-t`
- [ ] TUI mode for `-a`
- [ ] Follow mode for CLI `-f [blocks|transactions]`
