# Blockchain Challenge

Welcome the Energi Blockchain Challenge.

The instructions for the task are in #1.

To communicate with a reviewer, make sure to open a merge request and ask questions.

Your access to this repo will be rescinded in the _next 72 hours_.

Good Luck!


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



