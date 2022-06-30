package utils

var RpcUrl string = "https://nodeapi.test.energi.network/v1/jsonrpc"
var WsUrl string = "wss://nodeapi.test.energi.network/ws"
var BidContractAddress string = "0xBC452a95d170d17cE04F0960872127F2AB8181a2"
var ChainId = 49797
var Functions = []string{
	"function start()",
	"function checkGameRunning() view returns (bool)",
	"function totalWinnerAmount() view returns (uint256)",
}
