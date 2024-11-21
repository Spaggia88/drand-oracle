package config

type Config struct {
	DrandURLs             []string `envconfig:"DRAND_URLS" required:"true"`
	ChainHash             string   `envconfig:"CHAIN_HASH" required:"true"`
	DrandOracleAddress    string   `envconfig:"DRAND_ORACLE_ADDRESS" required:"true"`
	RPC                   string   `envconfig:"RPC" required:"true"`
	ChainID               int64    `envconfig:"CHAIN_ID" required:"true"`
	SetRandomnessGasLimit uint64   `envconfig:"SET_RANDOMNESS_GAS_LIMIT" required:"true"`
	SignerPrivateKey      string   `envconfig:"SIGNER_PRIVATE_KEY" required:"true"`
	SenderPrivateKey      string   `envconfig:"SENDER_PRIVATE_KEY" required:"true"`
	GenesisRound          uint64   `envconfig:"GENESIS_ROUND" required:"true"`
	MetricsPort           int      `envconfig:"METRICS_PORT" default:"4014"`
}
