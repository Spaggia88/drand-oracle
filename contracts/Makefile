# Directory configurations
CONTRACTS_DIR = src
OUT_DIR = ../updater/binding
ABI_DIR = out

# Contract names
CONTRACT = DrandOracle
INTERFACE = IDrandOracle

# File paths
CONTRACT_JSON = $(ABI_DIR)/$(CONTRACT).sol/$(CONTRACT).json
CONTRACT_ABI = $(ABI_DIR)/$(CONTRACT).sol/$(CONTRACT).abi
GO_OUT = $(OUT_DIR)/$(shell echo $(CONTRACT) | tr '[:upper:]' '[:lower:]').go

.PHONY: all clean generate-go

all: generate-go

# Ensure output directory exists
$(OUT_DIR):
	mkdir -p $(OUT_DIR)

# Generate Go bindings
generate-go: $(OUT_DIR)
	forge build
	forge inspect $(CONTRACTS_DIR)/$(CONTRACT).sol:$(CONTRACT) abi > $(CONTRACT_ABI)
	abigen --abi $(CONTRACT_ABI) \
		--pkg binding \
		--out $(GO_OUT)
	rm $(CONTRACT_ABI)

# Clean generated Go bindings
clean-go:
	rm -rf $(OUT_DIR)

# Deploy anvil using default anvil private key
deploy-anvil:
	PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
	OWNER=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
	SIGNER=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
	CHAIN_HASH=0x8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce \
	CHAIN=anvil \
	forge script ./script/deploy/DrandOracle.s.sol:DrandOracleScript --broadcast
