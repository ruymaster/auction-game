package controllers

import (
	"bidding/server/utils"
	"fmt"
	"math/big"
	"os"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func StartGame() {
	newWeb3, err := web3.NewWeb3(utils.RpcUrl)

	newWeb3.Eth.SetAccount(os.Getenv("OWNER_KEY"))
	newWeb3.Eth.SetChainId(int64(utils.ChainId))
	contract, err := newWeb3.Eth.NewContract(`[{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"start","inputs":[]}]`, utils.BidContractAddress)
	startInputData, err := contract.Methods("start").Inputs.Pack(nil)

	txHash, err := newWeb3.Eth.SyncSendRawTransaction(
		common.HexToAddress(utils.BidContractAddress),
		big.NewInt(0),
		300000,
		newWeb3.Utils.ToGWei(200),
		startInputData,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("---start transaction", txHash.TxHash)
}
