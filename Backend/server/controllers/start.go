package controllers

import (
	"bidding/server/utils"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func StartGame() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.Dial(utils.RpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("OWNER_KEY"))
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	biddingContractAddress := common.HexToAddress(utils.BidContractAddress)
	instance, err := NewBiddingwar(biddingContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(utils.ChainId)))
	if err != nil {
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
	}
	auth.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {

	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 300000
	auth.Value = big.NewInt(int64(0))

	tx, err := instance.Start(auth)

	totalWinnerAmount, err := instance.TotalWinnerAmount(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalWinnerAmount, fromAddress.String(), tx.Hash().String()) // "1.0"
}
