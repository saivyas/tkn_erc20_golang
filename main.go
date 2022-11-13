package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	GoToken "github.com/vyas-git/tkn_erc20_golang/go-gen"
)

var myenv map[string]string

const envLoc = ".env"

func main() {
	ctx := context.Background()
	client, err := ethclient.Dial(os.Getenv("GATEWAY")) //address of testnet
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	session := NewSession(ctx)

	totalSupply := new(big.Int)
	totalSupply.SetString("100000000000000000000000", 10)
	NewTokenDeploy(session, client, totalSupply)

}

func NewSession(ctx context.Context) (session GoToken.GoTokenSession) {
	pk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		log.Println("Error getting private key: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(5))

	if err != nil {
		log.Println("Error getting auth: ", err)
	}
	callOpts := bind.CallOpts{
		From:    auth.From,
		Context: ctx,
	}
	return GoToken.GoTokenSession{
		TransactOpts: *auth,
		CallOpts:     callOpts,
	}
}
func NewTokenDeploy(session GoToken.GoTokenSession, client *ethclient.Client, TotalSupply *big.Int) GoToken.GoTokenSession {
	log.Println("Creating new Token contract\n ")
	loadEnv()

	// deploy contract to get the contract addr, transaction object, adnd instance
	contractAddress, tx, instance, err := GoToken.DeployGoToken(&session.TransactOpts, client, TotalSupply)
	if err != nil {
		log.Fatalf("could not deploy token contract: %v\n", err)
	}

	//print the adress of the transaction - this can be used to up on ethresan the progress of the transaction
	fmt.Println("Contract deployed! Wait for tx %s to be confirmed.\n ", tx.Hash().Hex())

	session.Contract = instance

	//save the address of the deployed contract
	updateEnvFile("TOKEN_CONTRACTADDR", contractAddress.Hex())

	return session
}
func loadEnv() {
	var err error
	if err = godotenv.Load(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}

	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not read in env from %s: %s", envLoc, err)
	}
}
func updateEnvFile(key, val string) {
	myenv[key] = val
	if err := godotenv.Write(myenv, envLoc); err != nil {
		log.Printf("failed to update %s: %s", envLoc, err)
	}
}
