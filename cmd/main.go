package main

import (
	"bufio"
	"evmos-test/models"
	"evmos-test/pkgs/config"
	"evmos-test/services"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	config := config.NewConfig()

	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Failed to load environment variables! : " + err.Error())
	}

	url := config.EvmosUrl

	contractCounts := map[string]int{}
	walletAddresses := map[string]bool{}

	for i := 100; i <= 200; i++ {

		block, err := services.GetBlock(i, url)
		if err != nil {
			continue
		}

		m := min(10, len(block.Transactions))

		for _, transaction := range block.Transactions[:m] {

			receipt, err := services.GetTransactionReceipt(transaction.Hash, url)
			if err != nil {
				continue
			}

			contractAddress, isValid := receipt.ContractAddress.(string)
			if isValid {
				contractCounts[contractAddress] += 1
			}

			walletAddresses[transaction.To] = true
			walletAddresses[transaction.From] = true
		}
	}

	contracts := []models.Contract{}

	for address, count := range contractCounts {
		contracts = append(contracts, models.Contract{
			Address: address,
			Count:   count,
		})
	}

	sort.Slice(contracts, func(i, j int) bool {
		if contracts[i].Count > contracts[j].Count {
			return true
		} else if contracts[i].Count == contracts[j].Count && contracts[i].Address < contracts[j].Address {
			return true
		}
		return false
	})

	wallets := []models.Wallet{}

	for walletAddress := range walletAddresses {
		balance, err := services.GetBalance(walletAddress, url)
		if err != nil {
			continue
		}

		wallets = append(wallets, models.Wallet{
			Address: walletAddress,
			Balance: balance,
		})
	}

	sort.Slice(wallets, func(i, j int) bool {
		if wallets[i].Balance.Cmp(wallets[j].Balance) == 1 {
			return true
		} else if wallets[i].Balance.Cmp(wallets[j].Balance) == 0 && wallets[i].Address < wallets[j].Address {
			return true
		}
		return false
	})

	fileName := "example.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("contracts\n")
	if err != nil {
		fmt.Println("Error writing line to file:", err)
		return
	}

	for _, contract := range contracts {
		_, err := writer.WriteString(fmt.Sprintf("%s - %d\n", contract.Address, contract.Count))
		if err != nil {
			fmt.Println("Error writing line to file:", err)
			return
		}
	}

	_, err = writer.WriteString("wallets\n")
	if err != nil {
		fmt.Println("Error writing line to file:", err)
		return
	}

	for _, wallet := range wallets {
		_, err := writer.WriteString(fmt.Sprintf("%s - %s\n", wallet.Address, wallet.Balance.String()))
		if err != nil {
			fmt.Println("Error writing line to file:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}

	fmt.Println("File written successfully!")
}
