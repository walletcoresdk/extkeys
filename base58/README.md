base58
==========

[![Build Status](http://img.shields.io/travis/btcsuite/btcutil.svg)](https://travis-ci.org/btcsuite/btcutil)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/btcsuite/btcd/btcutil/base58)

Package base58 provides an API for encoding and decoding to and from the
modified base58 encoding.  It also provides an API to do Base58Check encoding,
as described [here](https://en.bitcoin.it/wiki/Base58Check_encoding).

A comprehensive suite of tests is provided to ensure proper functionality.

## Installation and Updating

```bash
$ go get -u github.com/btcsuite/btcd/btcutil/base58
```

## Examples

* [Decode Example](http://godoc.org/github.com/btcsuite/btcd/btcutil/base58#example-Decode)  
  Demonstrates how to decode modified base58 encoded data.
* [Encode Example](http://godoc.org/github.com/btcsuite/btcd/btcutil/base58#example-Encode)  
  Demonstrates how to encode data using the modified base58 encoding scheme.
* [CheckDecode Example](http://godoc.org/github.com/btcsuite/btcd/btcutil/base58#example-CheckDecode)  
  Demonstrates how to decode Base58Check encoded data.
* [CheckEncode Example](http://godoc.org/github.com/btcsuite/btcd/btcutil/base58#example-CheckEncode)  
  Demonstrates how to encode data using the Base58Check encoding scheme.

## License

Package base58 is licensed under the [copyfree](http://copyfree.org) ISC
License.



aa24ce23e2c87da7d7ef9f7816b8873a9db1043b4fa4f65b934f7764243b097c 私钥解析情况:

格式私钥(1)
私钥类型：	Hex格式
私钥Hex形式：	aa24ce23e2c87da7d7ef9f7816b8873a9db1043b4fa4f65b934f7764243b097c

生成公钥(2)
生成公钥(x,y)：	04dc52e055fe0258f82006bf82a6c1d473265f251629d699d2a3b7f16133dec378a5f858cb8389ec2a68db2fd5f03f8212ac6bb1f21151349cd9dcd53f547d65e7
压缩公钥(x)：	03dc52e055fe0258f82006bf82a6c1d473265f251629d699d2a3b7f16133dec378

生成Hash160(3)
压缩公钥hash160：	0dab4c763af10c9c1c29cd8e2763446450b662a9
未压缩公钥hash160：	c5057fb6a91f8b2ad2b743350db17ef4b0df03ad

生成钱包地址(4)
钱包地址：	12FGzY1e447Wifi4rxGEY6Ga9n18CTPiwh
未压缩钱包地址：	1JxkhrrV9Do7UNtR94UQCMoi1PPLyRapDu
钱包地址(P2SH)：	32wHv5W5bxRtoqQVz3vpxidWJJHqk1RjTT
未压缩钱包地址(P2SH)：	3KemdQLvh87VZYarGA8zczAe9ug4TnAMus
测试网络钱包地址：	mgmEHb6cs5YmVnBgaXEcN1Uu1mbq5RULXU
未压缩测试网络钱包地址：	myUhzuwTxFENFVN2rdSn2H22sNz3sDLTie


0x04dc52e055fe0258f82006bf82a6c1d473265f251629d699d2a3b7f16133dec378a5f858cb8389ec2a68db2fd5f03f8212ac6bb1f21151349cd9dcd53f547d65e7