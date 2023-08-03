/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"erc20Deployer/utils"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new instance of the ERC20 token smart contract on chain. Returns contract address.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		deploy()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}

func deploy() {

	config := utils.LoadConfiguration("./config.json")

	// Connect to our custom RPC
	conn, err := ethclient.Dial(config.RPCAddress)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum RPC: %v", err)
	}

	// Unlock the wallet
	fmt.Print("Wallet Password:")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	pass := string(bytepw)
	fmt.Print("\n")

	// Init a transactor for our chain
	// Give the wallet password to sign the tx
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(config.KeyStore), pass, big.NewInt(config.ChainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set the gas price
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	//
	fmt.Printf("Are you sure to deploy this ERC20 token contract to the current network (chain ID: %v)  [y/N]", config.ChainId)
	if askForConfirmation() {
		// Deploy the contract passing the newly created `auth` and `conn` vars
		address, tx, _, err := utils.DeployERC20Token(auth, conn)

		if err != nil {
			log.Fatalf("Failed to deploy new task contract: %v", err)
		}
		fmt.Printf("Contract pending deploy: 0x%x\n", address)
		fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

		time.Sleep(10 * time.Second) // Allow it to be processed by the local node :P
	} else {
		fmt.Println("Bye.")
	}

}

func askForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	case "":
		return askForConfirmation()
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter:")
		return askForConfirmation()
	}
}
