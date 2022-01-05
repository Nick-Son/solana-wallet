package cmd

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
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

func ImportOldWallet(RPCEndpoint string) (Wallet, error) {
	contents, _ := ioutil.ReadFile("key_data")
	privateKey := []byte(string(contents))
	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func GetBalance() (uint64, error) {
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
	balance, err := wallet.c.GetBalance(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
	)

	if err != nil {
		return 0, nil
	}

	return balance, nil
}

func RequestAirdrop(amount uint64) (string, error) {
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
	amount = amount * 1e9
	txhash, err := wallet.c.RequestAirdrop(
		context.TODO(),
		wallet.account.PublicKey.ToBase58(),
		amount,
	)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

func Transfer(receiver string, amount uint64) (string, error) {
	wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
	response, err := wallet.c.GetRecentBlockhash(context.TODO())

	if err != nil {
		return "", err
	}

	amount = amount * 1e9
	message := types.NewMessage(
		wallet.account.PublicKey,
		[]types.Instruction{
			sysprog.Transfer(
				wallet.account.PublicKey,
				common.PublicKeyFromString(receiver),
				amount,
			),
		},
		response.Blockhash,
	)

	tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account})
	if err != nil {
		return "", err
	}

	txhash, err := wallet.c.SendTransaction2(context.TODO(), tx)
	if err != nil {
		return "", err
	}

	return txhash, nil
}
