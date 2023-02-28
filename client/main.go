package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintEvents(infura *ethclient.Client, contractAddress common.Address) {
	beginBlock := int64(8571000)
	endBlock := int64(8572000)

	records, err := infura.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(beginBlock),
		ToBlock:   big.NewInt(endBlock),
		Addresses: []common.Address{contractAddress},
	})
	CheckErr(err)

	parsedAbi, err := abi.JSON(strings.NewReader(string(MainABI)))

	for _, rec := range records {
		fmt.Println(rec.TxHash.Hex())
		fmt.Println(rec.Data)

		event, err := parsedAbi.Unpack("addActor", rec.Data)
		CheckErr(err)

		fmt.Println(event)
	}
}

func PrepareTxn(infura *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	nonce, err := infura.PendingNonceAt(context.Background(), fromAddress)
	CheckErr(err)

	gasPrice, err := infura.SuggestGasPrice(context.Background())
	CheckErr(err)

	txn := bind.NewKeyedTransactor(privateKey)
	txn.Nonce = big.NewInt(int64(nonce))
	txn.Value = big.NewInt(0)
	txn.GasLimit = uint64(400000)
	txn.GasPrice = gasPrice

	return txn
}

func PlaceCreateActor(infura *ethclient.Client, contractAddress common.Address, fromAddress common.Address, privateKey *ecdsa.PrivateKey) {
	fmt.Println("Call PlaceCreateActor")

	txn := PrepareTxn(infura, fromAddress, privateKey)

	instance, err := NewMain(contractAddress, infura)
	CheckErr(err)

	result, err := instance.AddActor(txn, "id1", big.NewInt(1), big.NewInt(2), false)
	CheckErr(err)

	fmt.Printf("Placed event with transaction hash = %s\n", result.Hash().Hex())
}

func PlaceDeleteActor(infura *ethclient.Client, contractAddress common.Address, fromAddress common.Address, privateKey *ecdsa.PrivateKey) {
	fmt.Println("Call PlaceDeleteActor")

	txn := PrepareTxn(infura, fromAddress, privateKey)

	instance, err := NewMain(contractAddress, infura)
	CheckErr(err)

	result, err := instance.DelActor(txn, "id1")
	CheckErr(err)

	fmt.Printf("Placed event with transaction hash = %s\n", result.Hash().Hex())
}

func main() {
	privateKeyHex := os.Getenv("PRIVATE_KEY")

	infura, err := ethclient.Dial("https://goerli.infura.io/v3/03fb29f2a3134c3e913b5040738f9dab")
	CheckErr(err)

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	CheckErr(err)

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	contractAddress := common.HexToAddress("0xf0b408d7418a07C37FE90aFeA69406aeE8Cc9837")

	PlaceCreateActor(infura, contractAddress, fromAddress, privateKey)
	PlaceDeleteActor(infura, contractAddress, fromAddress, privateKey)
	PrintEvents(infura, contractAddress)
}
