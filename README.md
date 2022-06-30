## Game Mechanics

1. The bidding starts with 0 NRG and 60 minutes countdown.
2. Every next bid has to be higher than the previous one.
3. For each bid, the time extends for 10 minutes.
4. When the time runs out, the winner is the last bidder.
5. The reward is all the funds that were accumulated.


# Instructions to deploy smart contract on energi testnet
- cd SmartContracts
- cp .env.example .env (copy private key to deploy smart contract)
- nano .env
- yarn install
- yarn deploy --network energitest
# Testing Smart contract
- yarn bid-test
# Instructions to run backend
- cd Backend/server
- cp .env.example .env (copy owner key to run owner function of smart contract)
- go run .

# Starting Game with api
- curl -X POST http://localhost:8080/admin/start \
	  	-H 'authorization: Basic YWRtaW46MTIz' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"admin"}'
		
# Once game is started, backend will restart game automatically



