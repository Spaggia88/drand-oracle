ANVIL_RPC="http://localhost:8545"
ANVIL_DRAND_ORACLE_ADDRESS="0xa6F9208a7F9d8CFBDe24417CfAE3092AF31029d0"
ANVIL_SIGNER_PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
ANVIL_SENDER_PRIVATE_KEY="0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
ANVIL_CHAIN_ID=31337
ANVIL_GENESIS_ROUND=4496672

build:
	go build -o updater ./cmd/main.go

.PHONY: lint lint-fix
lint:
	@echo "--> Running linter"
	@golangci-lint run

lint-fix:
	@echo "--> Running linter auto fix"
	@golangci-lint run --fix

# Run updater locally
local-run-anvil:
	@echo "⚠️  WARNING: Please ensure that:"
	@echo "   1. Anvil is running (local Ethereum testnet)"
	@echo "   2. The Drand Oracle contract is deployed to Anvil using the `contracts/Makefile deploy-anvil` command"
	@echo "\nRunning updater against contract on anvil..."
	export DRAND_URLS="https://api.drand.sh,https://drand.cloudflare.com" && \
	export CHAIN_HASH="8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce" && \
	export DRAND_ORACLE_ADDRESS=$(ANVIL_DRAND_ORACLE_ADDRESS) && \
	export RPC=$(ANVIL_RPC) && \
	export CHAIN_ID=$(ANVIL_CHAIN_ID) && \
	export SET_RANDOMNESS_GAS_LIMIT=500000 && \
	export SIGNER_PRIVATE_KEY=$(ANVIL_SIGNER_PRIVATE_KEY) && \
	export SENDER_PRIVATE_KEY=$(ANVIL_SENDER_PRIVATE_KEY) && \
	export GENESIS_ROUND=$(ANVIL_GENESIS_ROUND) && \
	export MAX_RETRIES=2 && \
	go run --mod=mod ./cmd/main.go
