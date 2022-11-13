package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	GoToken "github.com/vyas-git/tkn_erc20_golang/go-gen"
)

func main() {
	ctx := context.Background()
	session := NewSession(ctx)

}

func NewSession(ctx context.Context) (session GoToken.GoTokenSession) {
	var pk *ecdsa.PrivateKey

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
func NewTokenDeploy() {

}
