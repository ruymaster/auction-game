package controllers

import (
	"bidding/server/utils"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// function to get owner auth option data for owner transaction
func getOwnerAuthOpts(client *ethclient.Client) *bind.TransactOpts {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("OWNER_KEY"))
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
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
	return auth
}

func processingAfterExited() {
	fmt.Println("---exit game")
	client, err := ethclient.Dial(utils.RpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	biddingContractAddress := common.HexToAddress(utils.BidContractAddress)
	instance, err := NewBiddingwar(biddingContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	res, err := instance.TotalWinnerAmount(nil)
	log.Println(" Total Winner Amount: ", res.Int64())
	if res.Int64() > 0 {
		rewardTx, err := instance.RewardToWinner(getOwnerAuthOpts(client))
		log.Println("-- reward transaction:", rewardTx.Hash())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		go time.AfterFunc(time.Duration(1)*time.Second, StartGame)
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

	biddingContractAddress := common.HexToAddress(utils.BidContractAddress)
	instance, err := NewBiddingwar(biddingContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.Start(getOwnerAuthOpts(client))
	fmt.Println("--start tx:", tx.Hash().String())
}

func CronEvent() {
	log.Println("--- Cron job to subscribe event from smart contract is strating...")
	client, err := ethclient.DialContext(context.Background(), utils.WsUrl)
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(utils.BidContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	biddingContract, err := NewBiddingwar(contractAddress, client)
	logs := make(chan types.Log)

	if err != nil {
		log.Fatal(err)
	}
	for {
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Fatal(err)
		}

		select {
		case err := <-sub.Err():
			if strings.Contains(err.Error(), "close 1006") {
				continue
			}
			log.Println("error from listener:", err)
		case vLog := <-logs:
			started, err := biddingContract.ParseGameStarted(vLog)
			if err != nil {
				bidInfo, errStarted := biddingContract.ParseBid(vLog)
				if errStarted != nil {

				} else {
					log.Println("Bid:", bidInfo.BidAmount, bidInfo.Bidder, bidInfo.BidTime)
				}
			} else {
				log.Printf("Game started at", started.StartTime.Int64(), started.GameIndex.Int64())
				go time.AfterFunc(time.Duration(181)*time.Second, processingAfterExited)
			}

		}
	}
}
