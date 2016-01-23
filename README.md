# be (blockchain explorer)

A Bitcoin Blockchain explorer for the terminal.

**be** uses the [Insight API](https://insight.is/) to get data from the Blockchain. By default it uses the Insight Live URL but in case you're run ning your own node, the URL can be defined by `-u <url>`

## Examples

Get Address info

```
be -a 1xsb94c9AMkj8GzhqYEJkieCXBpCZPvaF
```

Inspect Inputs and Outputs of a Transaction

```
be -t 4b8f9d2f4a53eb2c8de6ea54ef538edef19031792c4e1bcf003bea4c01d3a1da
```

Inspect Block

```
be -b 0000000000000000016491fbaabd1db0a89a4a281c4a14c27e3ec56f20c8b2bb
```
