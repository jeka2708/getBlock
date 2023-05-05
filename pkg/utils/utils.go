package utils

import (
	"log"
	"math/big"

	"getBlockTest/pkg/models"
)

func GetMaxValue(transactions []models.Transaction) models.Transaction {
	maxTransaction := transactions[0]
	var maxValue big.Int
	if _, ok := maxValue.SetString(transactions[0].Value[2:], 16); !ok {
		log.Fatal("Unable to parse value")
	}
	for _, transaction := range transactions {
		var val big.Int
		if _, ok := val.SetString(transaction.Value[2:], 16); !ok {
			log.Fatal("Unable to parse value")
		}
		if val.Cmp(&maxValue) == 1 {
			maxTransaction = transaction
			maxValue = val
		}
	}

	return maxTransaction
}
