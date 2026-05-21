package listener

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	encryptionAbe "eth-event-listener/internal/abe"
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
	"github.com/fentec-project/bn256"
	"github.com/fentec-project/gofe/abe"
	"github.com/fentec-project/gofe/data"
	"github.com/joho/godotenv"
)

type TransferEvent struct {
	TokenId         *big.Int       // Corresponds to 'uint tokenId'
	Sender          common.Address // Corresponds to 'address sender'
	AddressInChainB common.Address
	TokenURIValue   string
}

type PolicyAttribCheckResult struct {
	Used       []string // уникальные атрибуты из MSP
	Unknown    []string // атрибуты, которых нет в allowlist компании
	InvalidFmt []string // атрибуты, не прошедшие проверку формата (если включено)
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

			// before, after, _ := strings.Cut(event.TokenURIValue, "?")

			// encData, err := encryptionAbe.UnmarshalEncryptedDataFromString(before)
			// if err != nil {
			// 	log.Println("err unmarshaling tokenURI: ", err)
			// }
			sendToReceiverChain(event)
			// CheckAndSendToReceiverChain(encData, event, event.AddressInChainB.Hex())
		}
	}
}

func sendToReceiverChain(event TransferEvent) {
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

func checkAndSendToReceiverChain(encData encryptionAbe.EncryptedData, event TransferEvent, reciever string) {
	ok, err := checkRecipient(encData.Cipher.Msp, reciever)
	if err != nil {
		log.Println("error checking recipient:", err)
		fmt.Println("BLOCK")
		return
	}
	if ok {
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
	} else {
		fmt.Println("BLOCK wNFT MINT")
	}
}

func checkRecipient(msp *abe.MSP, recipientAddress string) (bool, error) {
	if msp == nil {
		return false, fmt.Errorf("nil msp")
	}
	if len(msp.RowToAttrib) == 0 {
		return false, fmt.Errorf("empty RowToAttrib")
	}
	if len(msp.Mat) != len(msp.RowToAttrib) {
		return false, fmt.Errorf("Mat/RowToAttrib length mismatch")
	}

	recipient, err := loadRecipientAttrs("", recipientAddress)
	if err != nil {
		return false, fmt.Errorf("load recipient attrs: %w", err)
	}

	var rows data.Matrix
	for i, attr := range msp.RowToAttrib {
		if recipient[attr] {
			rows = append(rows, append(data.Vector{}, msp.Mat[i]...))
		}
	}
	if len(rows) == 0 {
		return false, nil
	}

	p := bn256.Order
	cols := len(rows[0])

	aug := make([][]*big.Int, cols)
	for j := 0; j < cols; j++ {
		row := make([]*big.Int, len(rows)+1)
		for i := 0; i < len(rows); i++ {
			row[i] = new(big.Int).Mod(rows[i][j], p)
		}
		if j == 0 {
			row[len(rows)] = big.NewInt(1)
		} else {
			row[len(rows)] = big.NewInt(0)
		}
		aug[j] = row
	}

	return solvableModP(aug, len(rows), p), nil
}

func loadRecipientAttrs(envPath, address string) (map[string]bool, error) {
	if envPath == "" {
		envPath = ".env"
	}
	if err := godotenv.Load(envPath); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load %s: %w", envPath, err)
	}

	raw := strings.TrimSpace(os.Getenv("RECIPIENTS_BY_ADDRESS"))
	if raw == "" {
		return nil, fmt.Errorf("env %s is not set", "RECIPIENTS_BY_ADDRESS")
	}

	var dict map[string]string
	if err := json.Unmarshal([]byte(raw), &dict); err != nil {
		return nil, fmt.Errorf("parse %s as JSON: %w", "RECIPIENTS_BY_ADDRESS", err)
	}

	value, ok := lookupAddress(dict, address)
	if !ok {
		return nil, fmt.Errorf("address %q not found in %s", address, "RECIPIENTS_BY_ADDRESS")
	}

	dep, role, err := parseDepRole(value)
	if err != nil {
		return nil, fmt.Errorf("value for %s: %w", address, err)
	}

	attrs, err := encryptionAbe.GenerateAttributes(dep, role)
	if err != nil {
		return nil, fmt.Errorf("generateAttributes(%d,%d): %w", dep, role, err)
	}

	set := make(map[string]bool, len(attrs))
	for _, a := range attrs {
		set[a] = true
	}
	return set, nil
}

func solvableModP(aug [][]*big.Int, n int, p *big.Int) bool {
	m := len(aug)
	row := 0
	for col := 0; col < n && row < m; col++ {
		pivot := -1
		for r := row; r < m; r++ {
			if aug[r][col].Sign() != 0 {
				pivot = r
				break
			}
		}
		if pivot == -1 {
			continue
		}
		aug[row], aug[pivot] = aug[pivot], aug[row]

		inv := new(big.Int).ModInverse(aug[row][col], p)
		if inv == nil {
			return false
		}
		for c := col; c <= n; c++ {
			aug[row][c].Mul(aug[row][c], inv).Mod(aug[row][c], p)
		}

		for r := 0; r < m; r++ {
			if r == row || aug[r][col].Sign() == 0 {
				continue
			}
			factor := new(big.Int).Set(aug[r][col])
			for c := col; c <= n; c++ {
				t := new(big.Int).Mul(factor, aug[row][c])
				aug[r][c].Sub(aug[r][c], t).Mod(aug[r][c], p)
			}
		}
		row++
	}

	for r := 0; r < m; r++ {
		allZero := true
		for c := 0; c < n; c++ {
			if aug[r][c].Sign() != 0 {
				allZero = false
				break
			}
		}
		if allZero && aug[r][n].Sign() != 0 {
			return false
		}
	}
	return true
}

func lookupAddress(dict map[string]string, addr string) (string, bool) {
	if v, ok := dict[addr]; ok {
		return v, true
	}
	low := strings.ToLower(addr)
	for k, v := range dict {
		if strings.ToLower(k) == low {
			return v, true
		}
	}
	return "", false
}

func parseDepRole(s string) (int, int, error) {
	parts := strings.Split(strings.TrimSpace(s), ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("want 'dep,role', got %q", s)
	}
	dep, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("department: %w", err)
	}
	role, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("role: %w", err)
	}
	return dep, role, nil
}
