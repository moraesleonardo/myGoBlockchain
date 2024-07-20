package main

import (
	"fmt"
	"log"
	"github.com/moraesleonardo/myGoBlockchain/wallet"
	"github.com/moraesleonardo/myGoBlockchain/blockchain"
	
)

func init() {
	log.SetPrefix("Blockchain Node: ")
}

func main() {
	walletMiner := wallet.NewWallet()
	walletAlice := wallet.NewWallet()
	walletBob := wallet.NewWallet()

	// wallet transaction request
	t := wallet.NewTransaction(walletAlice.PrivateKey(), walletAlice.PublicKey(), walletAlice.BlockchainAddress(), walletBob.BlockchainAddress(), 23.0)

	// blockchain node transaction request handling
	blockchain := blockchain.NewBlockchain(walletMiner.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletAlice.BlockchainAddress(), walletBob.BlockchainAddress(), 23.0, walletAlice.PublicKey(), t.GenerateSignature())
	fmt.Println("Transaction added to transaction pool?", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("Miner has %.1f\n", blockchain.CalculateTotalAmount(walletMiner.BlockchainAddress()))
	fmt.Printf("Alice has %.1f\n", blockchain.CalculateTotalAmount(walletAlice.BlockchainAddress()))
	fmt.Printf("Bob has %.1f\n", blockchain.CalculateTotalAmount(walletBob.BlockchainAddress()))


}

