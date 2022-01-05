package cmd

import (
	"io/ioutil"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/types"
)

type Wallet struct {
	account types.Account  // Holds the wallet object
	c       *client.Client // Holds the RPC client object used to connect to Solana networks
}

func CreateNewWallet(RPCEndpoint string) Wallet {
	newAccount := types.NewAccount()      // create a new wallet using types.NewAccount()
	data := []byte(newAccount.PrivateKey) // Convert the private key to byte array

	err := ioutil.WriteFile("data", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return Wallet{
		newAccount,
		client.NewClient(RPCEndpoint),
	}
}
