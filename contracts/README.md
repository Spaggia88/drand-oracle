# 📄 Drand Oracle Smart Contract

## 🔎 Overview

The `DrandOracle` smart contract serves as an on-chain source of randomness by storing values from the drand network.

## 📚 Core Contracts

### 🎲 DrandOracle.sol

The main contract responsible for storing and managing randomness from the drand network.

#### 🛠️ Usage

- `getRandomnessFromRound(uint64 _round) external view returns (Random memory)`
  - Retrieves the complete randomness data for a specific round
  - Parameters:
    - `_round`: The round number to query
  - Returns:
    - `Random memory`: The Random struct containing the round's data
- `getRandomnessFromTimestamp(uint64 _timestamp) external view returns (Random memory)`
  - Retrieves the complete randomness data for a specific timestamp, or the latest timestamp prior to the given timestamp if no exact match is found
  - Parameters:
    - `_timestamp`: The timestamp to query
  - Returns:
    - `Random memory`: The Random struct containing the timestamp's data

### 🔒 Access Control

The contract implements a simple single-updater access control pattern:

- Only the designated updater address can submit new randomness
- Updater address is set during contract deployment
- Updater can be updated by the owner

### ⚠️ Limitations

1. The randomness stored in the contract is verifiable off-chain but not on-chain. This is due to BLS signatures used by Drand network are not yet supported on EVM.
2. Reliance on trusted updater.

### 🧪 Testing

Run tests with:

```bash
forge test
```

### 💻 Run Locally

1. Start local anvil chain

```bash
anvil
```

2. Deploy the contract

```bash
forge build
make deploy-anvil
```

### 🚀 Deploy

Edit the `.env` file to set the correct environment variables for the chain you want to deploy to. Then run:

```bash
FOUNDRY_PROFILE=deploy forge script ./script/deploy/DrandOracle.s.sol:DrandOracleScript --broadcast --verify --verifier blockscout
```

Note that for some testnets, `--gas-estimate-multiplier x` is required to avoid out-of-gas errors.

#### 🌍 Deployments

- Gravity Alpha Testnet Sepolia - [0x8d180A75fAc0De59BB1e833fc3838EE026a53aC5](https://explorer-sepolia.gravity.xyz/address/0x8d180A75fAc0De59BB1e833fc3838EE026a53aC5)
