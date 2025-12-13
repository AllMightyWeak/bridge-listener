package listener

import (
	"context"
	"os"

	"eth-event-listener/internal/conn"
	"eth-event-listener/internal/contract"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransferEvent struct {
	TokenId         *big.Int       // Corresponds to 'uint tokenId'
	Sender          common.Address // Corresponds to 'address sender'
	AddressInChainB common.Address
	TokenURIValue   string
}

var erc20Abi = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "addressInChainB",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tokenURIValue",
				"type": "string"
			}
		],
		"name": "TokenLocked",
		"type": "event"
	}]`

func EthereumListener() {
	client := conn.ConnectToWebSocket(os.Getenv("WEBSOCKET"))
	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("BRIDGE_ADDR")) // bridge address

	parsedABI, err := abi.JSON(strings.NewReader(erc20Abi))
	if err != nil {
		log.Fatal(err)
	}

	logs, sub := getEvents(client, contractAddress, parsedABI)
	defer sub.Unsubscribe()

	printData(sub, logs, parsedABI)

}

func getEvents(client *ethclient.Client, contractAddress common.Address, parsedABI abi.ABI) (chan types.Log, ethereum.Subscription) {
	eventSignatureHash := parsedABI.Events["TokenLocked"].ID

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{eventSignatureHash}}, // Optional: Filter by event signature
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to event logs: %v", err)
	}
	return logs, sub
}

func printData(sub ethereum.Subscription, logs chan types.Log, parsedABI abi.ABI) {
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)

		case vLog := <-logs:
			var event TransferEvent
			err := parsedABI.UnpackIntoInterface(&event, "TokenLocked", vLog.Data)
			if err != nil {
				log.Printf("Failed to unpack event data: %v", err)
				continue
			}

			fmt.Printf("--- New TokenLocked Event ---\n")
			log.Println()
			fmt.Printf("	BlockNumber: %d\n", vLog.BlockNumber)
			fmt.Printf("	TxHash: %s\n", vLog.TxHash.Hex())
			fmt.Printf("	Index: %d\n", vLog.Index)
			fmt.Printf("	Token ID: %s\n", event.TokenId.String())
			fmt.Printf("	Sender: %s\n", event.Sender.Hex())
			fmt.Printf("	Address in Chain B: %s\n", event.AddressInChainB.Hex())
			fmt.Printf("	Token URI: %s\n", event.TokenURIValue)

			client := conn.ConnectToWebSocket(os.Getenv("WEBSOCKET"))
			auth := conn.GetAccountAuth(client, os.Getenv("PRIVATE_KEY_BRIDGE")) // wallet private key

			conn, err := contract.NewWnft(common.HexToAddress(os.Getenv("WNFT_ADDR")), client) // nft contract address
			if err != nil {
				panic(err)
			}

			tx, err := conn.CreateToken(auth, common.HexToAddress(event.AddressInChainB.Hex()), event.TokenURIValue)
			if err != nil {
				log.Panic("error minting wrapped token: ", err)
			}
			fmt.Println("Hash: ", tx.Hash().Hex())
			fmt.Println("To: ", tx.To())
		}
	}
}
