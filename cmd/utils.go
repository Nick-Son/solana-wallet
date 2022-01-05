package cmd

import (
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/types"
)

type Wallet struct {
	// Holds the wallet object
	account types.Account
	// Holds the RPC client object used to connect to Solana networks
	c *client.Client
}
