# ERC20 Token Smart Contract
**ERC20 Token creation/deploy smartcontract in golang using openzepplin interface . Truffle & Hardhat alternative**

## Openzepplin Contracts
`git clone https://github.com/OpenZeppelin/openzeppelin-contracts.git
`


## Generate abi from solidity 

`solc --abi ./sol/GoToken.sol -o build 
`



## Generate bin from sol
 `solc --bin ./sol/GoToken.sol -o build
 `

## Generate Go Binding from abi 
`
abigen --bin=./build/GoToken.bin --abi=./build/GoToken.abi --pkg=GoToken --out=go-gen/GoToken.go
`