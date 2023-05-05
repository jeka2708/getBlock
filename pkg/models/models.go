package models

import "sync"

type BlockByNumberResponse struct {
	Addr         int64
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Value string `json:"value"`
	From  string `json:"from"`
	To    string `json:"to"`
}

type Result struct {
	sync.Mutex
	Transactions []Transaction
}
