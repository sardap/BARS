#!/bin/sh

geth --datadir . init ./myGenesis.json
geth --datadir . -rpc --rpcapi db,eth,net,web3,personal,web3 --rpccorsdomain "http://localhost:8000" --networkid 1114 console 2>> myEth.log
geth --datadir ./peer2DataDir --ipcdisable --networkid 1114 --port 30305 console 2>> myEth2.log


personal.unlockAccount(eth.accounts[0], "passphrase", 3000000)
web3.eth.defaultAccount = eth.accounts[0]
miner.setEtherbase(eth.accounts[0])


var abi = [{"constant":true,"inputs":[],"name":"returnTen","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"studentID","type":"uint32"},{"name":"documentId","type":"uint16"}],"name":"getValue","outputs":[{"name":"result","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"studentID","type":"uint32"},{"name":"documentId","type":"uint16"},{"name":"value","type":"string"}],"name":"add","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}];

var vaildcertsContract = web3.eth.contract(abi);
var vaildcertsContractInstance = vaildcertsContract.at("0x795cca35424e979b24d120d2b862430bf7207af8");
